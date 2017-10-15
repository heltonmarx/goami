package main

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/heltonmarx/goami/ami"
)

type Asterisk struct {
	socket *ami.Socket
	uuid   string

	events chan ami.Response
	stop   chan struct{}
	wg     sync.WaitGroup
}

// NewAsterisk initializes the AMI socket with a login and capturing the events.
func NewAsterisk(host string, username string, secret string) (*Asterisk, error) {
	socket, err := ami.NewSocket(host)
	if err != nil {
		return nil, err
	}
	uuid, err := ami.GetUUID()
	if err != nil {
		return nil, err
	}
	const events = "system,call,all,user"
	err = ami.Login(socket, username, secret, events, uuid)
	if err != nil {
		return nil, err
	}
	as := &Asterisk{
		socket: socket,
		uuid:   uuid,
		events: make(chan ami.Response),
		stop:   make(chan struct{}),
	}
	as.wg.Add(1)
	go as.run()
	return as, nil
}

// Logoff closes the current session with AMI.
func (as *Asterisk) Logoff() error {
	close(as.stop)
	as.wg.Wait()

	return ami.Logoff(as.socket, as.uuid)
}

// Events returns an channel with events received from AMI.
func (as *Asterisk) Events() <-chan ami.Response {
	return as.events
}

// SIPPeers fetch the list of SIP peers present on asterisk.
func (as *Asterisk) SIPPeers() error {
	resp, err := ami.SIPPeers(as.socket, as.uuid)
	switch {
	case err != nil:
		return err
	case len(resp) == 0:
		return errors.New("there's no sip peers configured")
	default:
		for _, v := range resp {
			message, err := ami.SIPShowPeer(as.socket, as.uuid, v.Get("ObjectName"))
			if err != nil {
				return err
			}
			fmt.Printf("message: %v\n", spew.Sdump(message))
		}
	}
	return nil
}

func (as *Asterisk) run() error {
	defer as.wg.Done()
	for {
		select {
		case <-as.stop:
			log.Printf("adios...\n")
			return nil
		default:
			events, err := ami.Events(as.socket)
			if err != nil {
				log.Printf("AMI events failed: %v\n", err)
				return err
			}
			as.events <- events
		}
	}
	return nil
}
