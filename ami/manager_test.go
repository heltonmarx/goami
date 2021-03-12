package ami

import (
	"context"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestLogin(t *testing.T) {
	var (
		actionID = "testid"
		user     = "testuser"
		secret   = "testsecret"
		events   = "Off"
		login    = "Action: Login\r\nActionID: testid\r\nUsername: testuser\r\nSecret: testsecret\r\nEvents: Off\r\n\r\n"
		response = "Asterisk Call Manager/1.0\r\nResponse: Success\r\nMessage: Authentication accepted\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, login, response)
	err := Login(ctx, client, user, secret, events, actionID)
	ensure.Nil(t, err)
}

func TestLogoff(t *testing.T) {
	var (
		actionID = "testid"
		logoff   = "Action: Logoff\r\nActionID: testid\r\n\r\n"
		response = "Response: Goodbye\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, logoff, response)
	err := Logoff(ctx, client, actionID)
	ensure.Nil(t, err)
}

func TestPing(t *testing.T) {
	var (
		actionID = "testid"
		ping     = "Action: Ping\r\nActionID: testid\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, ping, response)
	err := Ping(ctx, client, actionID)
	ensure.Nil(t, err)
}

func TestChallenge(t *testing.T) {
	var (
		actionID  = "testid"
		challenge = "Action: Challenge\r\nActionID: testid\r\nAuthType: MD5\r\n\r\n"
		response  = "Response: Success\r\nChallenge: 840415273\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, challenge, response)
	rsp, err := Challenge(ctx, client, actionID)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
	ensure.DeepEqual(t, rsp.Get("Challenge"), "840415273")
}

func TestCommand(t *testing.T) {
	var (
		actionID = "testid"
		cmd      = "Show Channels"
		input    = "Action: Command\r\nActionID: testid\r\nCommand: Show Channels\r\n\r\n"
		response = "Response: Follows\r\nChannel (Context Extension Pri ) State Appl. Data\r\n0 active channel(s)\r\n--END COMMAND--\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := Command(ctx, client, actionID, cmd)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Follows")
}

func TestCoreSettings(t *testing.T) {
	var (
		actionID = "testid"
		input    = "Action: CoreSettings\r\nActionID: testid\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := CoreSettings(ctx, client, actionID)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}

func TestCoreStatus(t *testing.T) {
	var (
		actionID = "testid"
		input    = "Action: CoreStatus\r\nActionID: testid\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := CoreStatus(ctx, client, actionID)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}

func TestCreateConfig(t *testing.T) {
	var (
		filename = "filename.txt"
		actionID = "testid"
		input    = "Action: CreateConfig\r\nActionID: testid\r\nFilename: filename.txt\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := CreateConfig(ctx, client, actionID, filename)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}

func TestDataGet(t *testing.T) {
	var (
		path     = "testpath"
		search   = "testsearch"
		filter   = "testfilter"
		actionID = "testid"
		input    = "Action: DataGet\r\nActionID: testid\r\nPath: testpath\r\nSearch: testsearch\r\nFilter: testfilter\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := DataGet(ctx, client, actionID, path, search, filter)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}

func TestEventFlow(t *testing.T) {
	var (
		eventMask = "off"
		actionID  = "testid"
		input     = "Action: Events\r\nActionID: testid\r\nEventMask: off\r\n\r\n"
		response  = "Response: Events Off\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := EventFlow(ctx, client, actionID, eventMask)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Events Off")
}

func TestGetConfigJSON(t *testing.T) {
	var (
		filename = "filename.txt"
		actionID = "testid"
		category = "category"
		filter   = "filter"
		input    = "Action: GetConfigJSON\r\nActionID: testid\r\nFilename: filename.txt\r\nCategory: category\r\nFilter: filter\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := GetConfigJSON(ctx, client, actionID, filename, category, filter)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}

func TestJabberSend(t *testing.T) {
	var (
		jabber   = "testjaber"
		jid      = "1234567890"
		message  = "foobar"
		actionID = "testid"
		input    = "Action: JabberSend\r\nActionID: testid\r\nJabber: testjaber\r\nJID: 1234567890\r\nMessage: foobar\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := JabberSend(ctx, client, actionID, jabber, jid, message)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}

func TestListCommands(t *testing.T) {
	var (
		actionID = "testid"
		input    = "Action: ListCommands\r\nActionID: testid\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := ListCommands(ctx, client, actionID)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}

func TestListCategories(t *testing.T) {
	var (
		filename = "testfile.txt"
		actionID = "testid"
		input    = "Action: ListCategories\r\nActionID: testid\r\nFilename: testfile.txt\r\n\r\n"
		response = "Response: Success\r\n\r\n"
	)
	ctx := context.Background()
	client := newClientMock(t, input, response)
	rsp, err := ListCategories(ctx, client, actionID, filename)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, rsp.Get("Response"), "Success")
}
