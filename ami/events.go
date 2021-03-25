package ami

/*
* El broker de eventos permite leer los eventos de asterisk y distribuirlos entre los subscriptores
* Esta implementación está basada en https://github.com/kljensen/golang-html5-sse-example/blob/master/server.go
 */

import (
	"context"
	_log "github.com/rs/zerolog/log"
	"sync"
)

var log = _log.With().Str("pkg", "ami").Str("context", "event-broker").Logger()

//NewBroker devuelve un nuevo EventsBroker de canales
func NewBroker() *EventsBroker {
	return &EventsBroker{
		clients:     make(map[chan Response]bool),
		subscribe:   make(chan (chan Response)),
		Unsubscribe: make(chan (chan Response)),
		Publish:     make(chan Response),
		stop:        make(chan struct{}),
		running:     false,
	}
}

// EventsBroker se encarga de leer los eventos AMI y distribuirlos entre
// los subscriptores
type EventsBroker struct {
	clients map[chan Response]bool

	subscribe   chan chan Response
	Unsubscribe chan chan Response
	Publish     chan Response
	mutex       sync.Mutex
	wg          sync.WaitGroup
	stop        chan struct{}
	running     bool
}

//Subscribe returns a new event channel
func (b *EventsBroker) Subscribe() chan Response {
	messageChan := make(chan Response, 10)
	b.subscribe <- messageChan
	return messageChan
}

// Start This EventsBroker method starts a new goroutine.  It handles
// the addition & removal of clients, as well as the broadcasting
// of messages out to clients that are currently attached.
func (b *EventsBroker) Start(ctx context.Context, client Client) {
	go func() {
		for {
			select {
			case <-b.stop:
				return
			case <-ctx.Done():
				return
			case s := <-b.subscribe:
				b.clients[s] = true
				b.wg.Add(1)

			case s := <-b.Unsubscribe:
				if b.clients[s] {
					b.mutex.Lock()
					delete(b.clients, s)
					close(s)
					b.mutex.Unlock()
					b.wg.Done()
				}
			case msg := <-b.Publish:
				b.mutex.Lock()
				// log.Trace().Interface("event", msg).Int("subscribers", len(b.clients)).Msg("Enviando evento a los suscriptores")
				for s := range b.clients {
					s <- msg
				}
				b.mutex.Unlock()
			}
		}
	}()

	go func() {
		for {
			select {
			default:
				if client.Connected() {
					events, err := read(ctx, client)
					if err != nil {
						log.Error().Err(err).Msg("Error al leer eventos asterisk")
						continue
					}
					b.Publish <- events
				}
			}
		}
	}()
	b.running = true

}

//Stop Detiene el broker de eventos
func (b *EventsBroker) Stop(ctx context.Context) {
	b.wg.Wait()
	b.running = false
	close(b.stop)

}

//IsRunning returns true if the event broker is running
func (b *EventsBroker) IsRunning() bool {
	return b.running
}
