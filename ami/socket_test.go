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

	"github.com/stretchr/testify/assert"
)

func TestSocketSend(t *testing.T) {
	const message = "Action: Login\r\nUsername: testuser\r\nSecret: testsecret\r\n\r\n"

	ctx, cancel := context.WithTimeout(t.Context(), time.Second)
	defer cancel()

	ln, err := net.Listen("tcp", ":")
	assert.NoError(t, err)
	defer ln.Close()

	incoming := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		conn, err := ln.Accept()
		assert.NoError(t, err)
		defer conn.Close()

		size := 0
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			assert.NoError(t, err)
			incoming <- msg
			size += len(msg)
			if size == len(message) {
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		var buffer bytes.Buffer
		for {
			select {
			case <-ctx.Done():
				t.Error("test timed out waiting for message: ", ctx.Err())
			case msg, ok := <-incoming:
				assert.True(t, ok)
				buffer.WriteString(msg)
				if strings.HasSuffix(buffer.String(), "\r\n\r\n") {
					assert.Equal(t, buffer.String(), message)
					return
				}
			}
		}
	}()
	socket, err := NewSocket(ctx, ln.Addr().String())
	assert.NoError(t, err)

	err = socket.Send(message)
	assert.NoError(t, err)

	wg.Wait()

	err = socket.Close(ctx)
	assert.NoError(t, err)
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
		assert.NoError(t, err)
		defer conn.Close()

		n, err := fmt.Fprintf(conn, response)
		assert.NoError(t, err)
		assert.True(t, n == len(response))
	}(ctx, ln)

	socket, err := NewSocket(ctx, ln.Addr().String())
	assert.NoError(t, err)

	rsp, err := socket.Recv(ctx)
	assert.NoError(t, err)
	assert.Equal(t, rsp, response)

	wg.Wait()
}
