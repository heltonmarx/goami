// Package pool manages AMI connections to multiple Asterisk servers at
// once: it dials, logs in, keeps each connection alive with automatic
// reconnect + backoff, and multiplexes every server's events onto a single
// channel. Built against github.com/heltonmarx/goami's context-based API
// (ami.NewSocket, ami.Connect, ami.Login, ami.Events, ami.Logoff).
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/heltonmarx/goami/ami"
)

// ServerConfig describes one Asterisk AMI endpoint.
type ServerConfig struct {
	Name     string // logical name, e.g. "sp-fortaleza", "sp-01"
	Host     string // host:port, e.g. "10.0.1.10:5038"
	Username string
	Secret   string
	Events   string // AMI event mask, e.g. "system,call,all,user"; "" -> default
}

// Event wraps an AMI event with the server it came from, so a single
// consumer can tell which PBX generated it.
type Event struct {
	Server string
	Data   ami.Response
}

// conn holds the mutable state for one server's connection.
type conn struct {
	cfg ServerConfig

	mu      sync.RWMutex
	socket  *ami.Socket
	uuid    string
	healthy bool

	refCount sync.WaitGroup
}

// Pool manages AMI connections to multiple Asterisk servers.
type Pool struct {
	conns map[string]*conn
	order []string // stable order, used for round-robin

	next   atomic.Uint64
	events chan Event

	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewPool builds a Pool for the given servers. Call Start to actually connect.
func NewPool(configs []ServerConfig) *Pool {
	p := &Pool{
		conns:  make(map[string]*conn, len(configs)),
		order:  make([]string, 0, len(configs)),
		events: make(chan Event, 256),
	}
	for _, c := range configs {
		p.conns[c.Name] = &conn{cfg: c}
		p.order = append(p.order, c.Name)
	}
	return p
}

// Events returns the channel where events from every connected server are
// published. Callers should keep draining it; a full buffer will make
// individual server goroutines block.
func (p *Pool) Events() <-chan Event {
	return p.events
}

// Start connects to every configured server in the background and keeps
// each one alive (with reconnect + backoff) until ctx is cancelled or
// Close is called.
func (p *Pool) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	p.cancel = cancel

	for _, name := range p.order {
		c := p.conns[name]
		p.wg.Add(1)
		go p.keepAlive(ctx, c)
	}
}

// Close stops every connection, waits for goroutines to exit, and closes
// the events channel. Safe to call once after Start.
func (p *Pool) Close() {
	if p.cancel != nil {
		p.cancel()
	}
	p.wg.Wait()

	// Drain any remaining events in the channel.
	for len(p.events) > 0 {
		<-p.events
	}
	close(p.events)
}

// keepAlive owns the full lifecycle for one server: connect, login, read
// events until the connection drops, then reconnect with exponential
// backoff. Runs until ctx is cancelled.
func (p *Pool) keepAlive(ctx context.Context, c *conn) {
	defer p.wg.Done()

	const maxBackoff = 30 * time.Second
	backoff := time.Second

	for ctx.Err() == nil {
		if err := p.connect(ctx, c); err != nil {
			log.Printf("ami[%s]: connect failed: %v (retry in %s)", c.cfg.Name, err, backoff)
			select {
			case <-ctx.Done():
				return
			case <-time.After(backoff):
			}
			if backoff < maxBackoff {
				backoff *= 2
			}
			continue
		}
		backoff = time.Second // reset once we get a clean connection

		p.readEvents(ctx, c) // blocks until the connection drops or ctx ends

		// Wait for any outstanding references before closing the socket.
		c.refCount.Wait()

		c.mu.Lock()
		c.healthy = false
		if c.socket != nil {
			_ = c.socket.Close(ctx)
			c.socket = nil
		}
		c.mu.Unlock()
	}
}

func (p *Pool) connect(ctx context.Context, c *conn) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	socket, err := ami.NewSocket(ctx, c.cfg.Host)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	if ok, err := ami.Connect(ctx, socket); err != nil || !ok {
		_ = socket.Close(ctx)
		return fmt.Errorf("handshake: %w", err)
	}

	uuid, err := ami.GetUUID()
	if err != nil {
		_ = socket.Close(ctx)
		return fmt.Errorf("uuid: %w", err)
	}

	events := c.cfg.Events
	if events == "" {
		events = "system,call,all,user"
	}
	if err := ami.Login(ctx, socket, c.cfg.Username, c.cfg.Secret, events, uuid); err != nil {
		_ = socket.Close(ctx)
		return fmt.Errorf("login: %w", err)
	}

	c.mu.Lock()
	c.socket = socket
	c.uuid = uuid
	c.healthy = true
	c.mu.Unlock()

	log.Printf("ami[%s]: connected to %s", c.cfg.Name, c.cfg.Host)
	return nil
}

func (p *Pool) readEvents(ctx context.Context, c *conn) {
	// Copy the socket pointer under lock; keepAlive guarantees the socket
	// stays alive while this goroutine runs.
	c.mu.RLock()
	socket := c.socket
	c.mu.RUnlock()

	if socket == nil {
		return
	}

	for {
		if ctx.Err() != nil {
			return
		}

		resp, err := ami.Events(ctx, socket)
		if err != nil {
			if ctx.Err() != nil {
				return // normal shutdown
			}
			log.Printf("ami[%s]: connection lost: %v", c.cfg.Name, err)
			return
		}

		select {
		case p.events <- Event{Server: c.cfg.Name, Data: resp}:
		case <-ctx.Done():
			return
		}
	}
}

// Get returns the live socket and ActionID for a specific server by name,
// or an error if that server isn't currently connected. Use this when an
// action (e.g. Originate) must target one particular PBX.
func (p *Pool) Get(name string) (socket *ami.Socket, actionID string, err error) {
	c, ok := p.conns[name]
	if !ok {
		return nil, "", fmt.Errorf("unknown server %q", name)
	}
	c.mu.RLock()
	if !c.healthy || c.socket == nil {
		c.mu.RUnlock()
		return nil, "", fmt.Errorf("server %q is not connected", name)
	}
	socket = c.socket
	actionID = c.uuid
	c.mu.RUnlock()

	// Increment reference count; caller must call Release(name) when done.
	c.refCount.Add(1)
	return socket, actionID, nil
}

// Next returns a healthy connection using round-robin selection. Useful
// when any server will do, e.g. distributing Originate calls across a
// fleet of otherwise-identical PBXs.
func (p *Pool) Next() (name string, socket *ami.Socket, actionID string, err error) {
	n := len(p.order)
	if n == 0 {
		return "", nil, "", fmt.Errorf("pool is empty")
	}
	for range n {
		idx := p.next.Add(1)
		idx = idx % uint64(n)
		candidate := p.order[idx]
		c := p.conns[candidate]

		c.mu.RLock()
		healthy, s, uuid := c.healthy, c.socket, c.uuid
		c.mu.RUnlock()

		if healthy {
			// Increment reference count; caller must call Release(name) when done.
			c.refCount.Add(1)
			return candidate, s, uuid, nil
		}
	}
	return "", nil, "", fmt.Errorf("no healthy servers available")
}

// Release decrements the reference count for a server's connection.
// Must be called after every successful Get or Next to allow the
// connection to be closed cleanly on shutdown.
//
// Calling Release twice for the same name (double-release) will cause
// a panic because the underlying WaitGroup counter becomes negative.
func (p *Pool) Release(name string) {
	c, ok := p.conns[name]
	if !ok {
		return
	}
	c.refCount.Done()
}

// Healthy reports which configured servers currently have a live
// connection. Handy for a /healthz endpoint or a Prometheus gauge.
//
// The returned map is a point-in-time snapshot; the actual health of
// a server may change immediately after the call returns.
func (p *Pool) Healthy() map[string]bool {
	status := make(map[string]bool, len(p.order))
	for _, name := range p.order {
		c := p.conns[name]
		c.mu.RLock()
		status[name] = c.healthy
		c.mu.RUnlock()
	}
	return status
}
