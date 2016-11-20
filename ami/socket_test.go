package ami

import (
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/facebookgo/ensure"
)

func TestSocketSend(t *testing.T) {
	message := "Action: Login\r\nUsername: testuser\r\nSecret: testsecret\r\n\r\n"

	done := make(chan struct{})
	srv := newServer(t, func(conn net.Conn) {
		defer conn.Close()
		buf, err := ioutil.ReadAll(conn)
		ensure.Nil(t, err)
		ensure.DeepEqual(t, message, string(buf))
		close(done)
	})
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
	response := "Asterisk Call Manager/1.0\r\nResponse: Success\r\nMessage: Authentication accepted\r\n\r\n"

	wait := make(chan struct{})
	srv := newServer(t, func(conn net.Conn) {
		defer conn.Close()
		n, err := conn.Write([]byte(response))
		ensure.Nil(t, err)
		ensure.True(t, n == len(response))
		<-wait
	})
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

func newServer(t *testing.T, handler func(net.Conn)) *testServer {
	ln, err := net.Listen("tcp", ":")
	ensure.Nil(t, err)

	ts := &testServer{
		ln:   ln,
		stop: make(chan struct{}),
	}
	go func() {
		for {
			select {
			case <-ts.stop:
				return
			case <-time.After(100 * time.Millisecond):
				conn, err := ts.ln.Accept()
				if err != nil {
					ensure.Nil(t, err)
					continue
				}
				go handler(conn)
			}
		}
	}()
	return ts
}

func (ts *testServer) Close() error {
	close(ts.stop)
	return ts.ln.Close()
}

func (ts *testServer) Addr() string {
	return ts.ln.Addr().String()
}
