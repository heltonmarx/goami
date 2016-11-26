package ami

// Client provides methods to client socket methods.
type Client interface {
	// Connected returns the client status.
	Connected() bool

	// Close closes the client connection.
	Close() error

	// Send sends data from client to server.
	Send(message string) error

	// Recv receives a string from server.
	Recv() (string, error)
}
