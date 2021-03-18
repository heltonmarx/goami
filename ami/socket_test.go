package ami

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net"
	"strings"
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

	incoming := make(chan string)
	done := make(chan struct{})

	go func() {
		conn, err := ln.Accept()
		ensure.Nil(t, err)
		defer conn.Close()

		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			ensure.Nil(t, err)
			incoming <- msg
		}
	}()

	go func() {
		var buffer bytes.Buffer
		for {
			select {
			case msg, ok := <-incoming:
				ensure.True(t, ok)
				buffer.WriteString(msg)
				if strings.HasSuffix(buffer.String(), "\r\n\r\n") {
					ensure.DeepEqual(t, buffer.String(), message)
					close(done)
				}
			}
		}
	}()
	socket, err := NewSocket(ctx, ln.Addr().String())
	ensure.Nil(t, err)

	err = socket.Send(message)
	ensure.Nil(t, err)

	<-done

	err = socket.Close(ctx)
	ensure.Nil(t, err)
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
