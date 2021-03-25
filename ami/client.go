package ami

import "context"

// Client defines an interface to client socket.
type Client interface {
	// Connected returns the client status.
	Connected() bool

	// Close closes the client connection.
	Close(ctx context.Context) error

	// Send sends data from client to server.
	Send(message string) error

	// Recv receives a string from server.
	Recv(ctx context.Context) (string, error)

	//NewEventChannel returns a new chan for receiving events
	Events() *EventsBroker
}
