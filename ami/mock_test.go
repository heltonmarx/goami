package ami

import (
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

func (m mockClient) Close() error {
	return m.closeFn()
}

func (m mockClient) Send(message string) error {
	return m.sendFn(message)
}

func (m mockClient) Recv() (string, error) {
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
			ensure.DeepEqual(t, message, input)
			return nil
		},
		recvFn: func() (string, error) {
			return output, nil
		},
	}
}
