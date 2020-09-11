package ami

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/facebookgo/ensure"
)

func TestSocketSend(t *testing.T) {
	const message = "Action: Login\r\nUsername: testuser\r\nSecret: testsecret\r\n\r\n"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ln, err := net.Listen("tcp", ":")
	ensure.Nil(t, err)
	defer ln.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context, ln net.Listener) {
		defer wg.Done()
		conn, err := ln.Accept()
		ensure.Nil(t, err)
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		ensure.Nil(t, err)

		msg := string(buf[:])
		ensure.DeepEqual(t, msg, message)
	}(ctx, ln)

	socket, err := NewSocket(ctx, ln.Addr().String())
	ensure.Nil(t, err)

	err = socket.Send(message)
	ensure.Nil(t, err)

	err = socket.Close(ctx)
	ensure.Nil(t, err)

	wg.Wait()
}

func TestSocketRecv(t *testing.T) {
	const response = "Asterisk Call Manager/1.0\r\n"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ln, err := net.Listen("tcp", ":")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context, ln net.Listener) {
		defer wg.Done()

		conn, err := ln.Accept()
		ensure.Nil(t, err)
		defer conn.Close()

		n, err := fmt.Fprintf(conn, response)
		ensure.Nil(t, err)
		ensure.True(t, n == len(response))
	}(ctx, ln)

	socket, err := NewSocket(ctx, ln.Addr().String())
	ensure.Nil(t, err)

	rsp, err := socket.Recv(ctx)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp, response)

	wg.Wait()
}
