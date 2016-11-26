package ami

import (
	"bufio"
	"bytes"
	"net"
	"strings"
)

// Socket holds the socket client connection data.
type Socket struct {
	conn net.Conn
}

// NewSocket provides a new socket client, connecting to a tcp server.
func NewSocket(address string) (*Socket, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &Socket{
		conn: conn,
	}, nil
}

// Connected returns the socket status, true for connected,
// false for disconnected.
func (s *Socket) Connected() bool {
	if s.conn == nil {
		return false
	}
	return true
}

// Close closes socket connection.
func (s *Socket) Close() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}

// Send sends data to socket using fprintf format.
func (s *Socket) Send(message string) error {
	_, err := s.conn.Write([]byte(message))
	return err
}

// Recv receives a string from socket server.
func (s *Socket) Recv() (string, error) {
	var buf []byte
	reader := bufio.NewReader(s.conn)
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			return "", err
		}
		buf = append(buf, b...)
		if (len(bytes.TrimSpace(b)) == 0 &&
			strings.Contains(string(buf), string('\n')) == true) || reader.Buffered() == 0 {
			break
		}
	}
	return string(buf), nil
}
