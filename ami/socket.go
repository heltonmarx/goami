package ami

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

// Socket holds the socket client connection data.
type Socket struct {
	conn     net.Conn
	incoming chan string
	shutdown chan struct{}
	wg       sync.WaitGroup
	events   *EventsBroker
}

// NewSocket provides a new socket client, connecting to a tcp server.
func NewSocket(ctx context.Context, address string) (*Socket, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	s := &Socket{
		conn:     conn,
		incoming: make(chan string, 32),
		shutdown: make(chan struct{}),
		events:   NewBroker(),
	}

	// s.events.Start(ctx, s)
	s.run(ctx, conn)
	//start the event broker
	return s, nil
}

// Connected returns the socket status, true for connected,
// false for disconnected.
func (s *Socket) Connected() bool {
	return s.conn != nil
}

// Close closes socket connection.
func (s *Socket) Close(ctx context.Context) error {
	close(s.shutdown)

	// wait for shutdown of run process
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.NewTimer(150 * time.Millisecond).C:
	case <-ctx.Done():
		return s.terminate()
	}
	return nil
}

// Send sends data to socket using fprintf format.
func (s *Socket) Send(message string) error {
	_, err := fmt.Fprintf(s.conn, message)
	return err
}

// Recv receives a string from socket server.
func (s *Socket) Recv(ctx context.Context) (string, error) {
	var buffer bytes.Buffer
	for {
		select {
		case msg, ok := <-s.incoming:
			if !ok {
				continue
			}
			buffer.WriteString(msg)
			if strings.HasSuffix(buffer.String(), "\r\n") {
				return buffer.String(), nil
			}
		case <-s.shutdown:
		case <-ctx.Done():
			return buffer.String(), io.EOF
		}
	}
}

func (s *Socket) run(ctx context.Context, conn net.Conn) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		reader := bufio.NewReader(conn)
		for {
			select {
			case <-s.shutdown:
			case <-ctx.Done():
				s.terminate()
				return
			default:
				msg, err := reader.ReadString('\n')
				if err != nil {
					return
				}
				s.incoming <- msg
			}
		}
	}()
}

func (s *Socket) terminate() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}

//Events returns the eventBroker
func (s *Socket) Events() *EventsBroker {
	return s.events
}
