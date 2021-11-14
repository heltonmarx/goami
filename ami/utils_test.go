package ami

import (
	"context"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestParseResponse(t *testing.T) {
	response := "Asterisk Call Manager/1.0\r\nResponse: Success\r\nMessage: Authentication accepted\r\n\r\n"
	rsp, err := parseResponse(response)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
	ensure.DeepEqual(t, rsp.Get("Message"), "Authentication accepted")
}

func TestSend(t *testing.T) {
	var (
		login    = "Action: Login\r\nActionID: testid\r\n\r\n"
		response = "Asterisk Call Manager/1.0\r\nResponse: Success\r\nMessage: Authentication accepted\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, login, response)
	rsp, err := send(ctx, client, "Login", "", nil)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
	ensure.DeepEqual(t, rsp.Get("Message"), "Authentication accepted")
}
