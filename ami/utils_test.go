package ami

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	response := "Asterisk Call Manager/1.0\r\nResponse: Success\r\nMessage: Authentication accepted\r\n\r\n"
	rsp, err := parseResponse(response)
	assert.NoError(t, err)
	assert.Equal(t, rsp.Get("Response"), "Success")
	assert.Equal(t, rsp.Get("Message"), "Authentication accepted")
}

func TestSend(t *testing.T) {
	var (
		login    = "Action: Login\r\nActionID: testid\r\n\r\n"
		response = "Asterisk Call Manager/1.0\r\nResponse: Success\r\nMessage: Authentication accepted\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, login, response)
	rsp, err := send(ctx, client, "Login", "testid", nil)
	assert.NoError(t, err)
	assert.Equal(t, rsp.Get("Response"), "Success")
	assert.Equal(t, rsp.Get("Message"), "Authentication accepted")
}
