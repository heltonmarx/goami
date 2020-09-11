package ami

import (
	"context"
	"strings"
	"testing"

	"github.com/facebookgo/ensure"
)

type mockClient struct {
	connectedFn func() bool
	closeFn     func() error
	sendFn      func(message string) error
	recvFn      func() (string, error)
}

func (m mockClient) Connected() bool {
	return m.connectedFn()
}

func (m mockClient) Close(ctx context.Context) error {
	return m.closeFn()
}

func (m mockClient) Send(message string) error {
	return m.sendFn(message)
}

func (m mockClient) Recv(ctx context.Context) (string, error) {
	return m.recvFn()
}

func newClientMock(t *testing.T, input, output string) mockClient {
	return mockClient{
		connectedFn: func() bool {
			return true
		},
		closeFn: func() error {
			return nil
		},
		sendFn: func(message string) error {
			verifyResponse(t, message, input)
			return nil
		},
		recvFn: func() (string, error) {
			return output, nil
		},
	}
}

func verifyResponse(t *testing.T, actual, expect string) {
	as := strings.Split(actual, "\r\n")
	es := strings.Split(expect, "\r\n")

	ensure.DeepEqual(t, len(as), len(es))
	for _, m := range as {
		found := false
		for _, n := range es {
			if m == n {
				found = true
			}
		}
		ensure.True(t, found)
	}
}
