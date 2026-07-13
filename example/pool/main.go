// Command poolexample connects to several Asterisk servers concurrently
// using the pool package, prints incoming events tagged by server, and
// shows how to target a single server for an action-style call.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/heltonmarx/goami/ami"
)

func main() {
	username := flag.String("username", "admin", "AMI username (shared by all servers in this example)")
	secret := flag.String("secret", "admin", "AMI secret (shared by all servers in this example)")
	flag.Parse()

	// In a real service these would come from config/env/service discovery,
	// each with its own credentials if needed.
	servers := []ServerConfig{
		{Name: "sp-01", Host: "10.0.1.10:5038", Username: *username, Secret: *secret},
		{Name: "sp-02", Host: "10.0.2.10:5038", Username: *username, Secret: *secret},
		{Name: "sp-03", Host: "10.0.3.10:5038", Username: *username, Secret: *secret},
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	p := NewPool(servers)
	p.Start(ctx)
	defer p.Close()

	// Consume events from every server on one channel.
	go func() {
		for evt := range p.Events() {
			log.Printf("[%s] event=%s channel=%s",
				evt.Server, evt.Data.Get("Event"), evt.Data.Get("Channel"))
		}
	}()

	// Example: target one specific server for an action.
	if socket, actionID, err := p.Get("sp-02"); err == nil {
		if resp, err := ami.CoreStatus(ctx, socket, actionID); err == nil {
			log.Printf("sp-02 core status: %s", resp.Get("CoreStartupDate"))
		}
		p.Release("sp-02")
	}

	// Example: don't care which server, just need any healthy one.
	if name, socket, actionID, err := p.Next(); err == nil {
		if resp, err := ami.CoreSettings(ctx, socket, actionID); err == nil {
			log.Printf("picked %s via round-robin, AMI version %s", name, resp.Get("AMIversion"))
		}
		p.Release(name)
	}

	<-ctx.Done()
	log.Println("shutting down")
}
