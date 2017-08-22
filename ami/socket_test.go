package ami

import (
	"context"
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/facebookgo/ensure"
)

func TestSocketSend(t *testing.T) {
	message := "Action: Login\r\nUsername: testuser\r\nSecret: testsecret\r\n\r\n"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan struct{})
	srv, err := newServer(ctx, func(conn net.Conn) {
		defer conn.Close()
		buf, err := ioutil.ReadAll(conn)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, message, string(buf))
		close(done)
	})
	ensure.Nil(t, err)
	defer srv.Close()

	socket, err := NewSocket(srv.Addr())
	ensure.Nil(t, err)

	err = socket.Send(message)
	ensure.Nil(t, err)

	err = socket.Close()
	ensure.Nil(t, err)

	<-done
}

func TestSocketRecv(t *testing.T) {
	response := "Asterisk Call Manager/1.0\r\n"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wait := make(chan struct{})
	srv, err := newServer(ctx, func(conn net.Conn) {
		defer conn.Close()
		n, err := conn.Write([]byte(response))
		ensure.Nil(t, err)
		ensure.True(t, n == len(response))
		<-wait
	})
	ensure.Nil(t, err)
	defer srv.Close()

	socket, err := NewSocket(srv.Addr())
	ensure.Nil(t, err)

	rsp, err := socket.Recv()
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp, response)
	close(wait)
}

// testServer used to server unit test connections.
type testServer struct {
	ln   net.Listener
	stop chan struct{}
}

func newServer(ctx context.Context, handler func(net.Conn)) (*testServer, error) {
	ln, err := net.Listen("tcp", ":")
	if err != nil {
		return nil, err
	}

	ts := &testServer{
		ln:   ln,
		stop: make(chan struct{}),
	}
	go func() {
		for {
			select {
			case <-ts.stop:
				return
			case <-ctx.Done():
				return
			default:
				conn, err := ts.ln.Accept()
				if err != nil {
					time.Sleep(10 * time.Millisecond)
					continue
				}
				go handler(conn)
			}
		}
	}()
	return ts, nil
}

func (ts *testServer) Close() error {
	close(ts.stop)
	return ts.ln.Close()
}

func (ts *testServer) Addr() string {
	return ts.ln.Addr().String()
}
