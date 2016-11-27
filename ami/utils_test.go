package ami

import (
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
	client := newClientMock(t, login, response)
	rsp, err := send(client, "Login", "testid", nil)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
	ensure.DeepEqual(t, rsp.Get("Message"), "Authentication accepted")
}

func TestRequestEvent(t *testing.T) {
	var (
		peerEntryList = []string{
			"Response: Success\r\nMessage: Peer status list will follow\r\n\r\n",
			"Event: PeerEntry\r\n" +
				"Channeltype: SIP\r\n" +
				"ObjectName: 9915057\r\n" +
				"ChanObjectType: peer\r\n" +
				"IPaddress: 10.64.72.166\r\n" +
				"IPport: 5060\r\n" +
				"Dynamic: yes\r\n" +
				"Natsupport: no\r\n" +
				"ACL: no\r\n" +
				"Status: OK (5 ms)\r\n\r\n",
			"Event: PeerlistComplete\r\nListItems: 205\r\n\r\n",
		}
		response = Response{
			"Event":          []string{"PeerEntry"},
			"Channeltype":    []string{"SIP"},
			"ObjectName":     []string{"9915057"},
			"ChanObjectType": []string{"peer"},
			"IPaddress":      []string{"10.64.72.166"},
			"IPport":         []string{"5060"},
			"Dynamic":        []string{"yes"},
			"Natsupport":     []string{"no"},
			"ACL":            []string{"no"},
			"Status":         []string{"OK (5 ms)"},
		}
	)
	rsp, finish, err := parseEvent("PeerEntry", "PeerlistComplete", peerEntryList)
	ensure.Nil(t, err)
	ensure.True(t, finish)
	ensure.DeepEqual(t, rsp[0], response)
}
