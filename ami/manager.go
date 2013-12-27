package ami

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func sendCmd(socket *Socket, cmd []string) error {
	for _, s := range cmd {
		err := socket.Send("%s", s)
		if err != nil {
			fmt.Printf("send login error:[%v]\n", err)
			return err
		}
	}
	return nil
}

func Connect(socket *Socket) (bool, error) {
	if !socket.Connected() {
		fmt.Printf("could not connect to AMI\n")
	}
	if answer, err := socket.Recv(); err != nil || !strings.Contains(answer, "Asterisk Call Manager") {
		return false, errors.New("AMI connection failed")
	}
	return true, nil
}

//  Login
//      Login Manager.
//
func Login(socket *Socket, user, secret, events, actionID string) (bool, error) {
	// verify parameters
	if len(user) == 0 || len(secret) == 0 || len(events) == 0 || len(actionID) == 0 {
		return false, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Login",
		"\r\nUsername: ",
		user,
		"\r\nSecret: ",
		secret,
		"\r\nEvents: ",
		events,
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	message, err := getMessage(socket, command, actionID)
	if err != nil {
		return false, err
	}

	if message["Response"] != "Success" {
		return false, errors.New(message["Message"])
	}
	return true, nil
}

//  Logoff
//      Logoff the current manager session.      
//
func Logoff(socket *Socket, actionID string) (bool, error) {
	// verify parameters
	if len(actionID) == 0 {
		return false, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Logoff",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	message, err := getMessage(socket, command, actionID)
	if err != nil {
		return false, err
	}
	if message["Response"] != "Goodbye" {
		return false, errors.New(message["Message"])
	}
	return true, nil
}

func GetUUID() (string, error) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Printf("open /dev/urandom error:[%v]\n", err)
		return "", err
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6],
		b[6:8], b[8:10], b[10:])
	return uuid, nil
}

//  Ping
//      A 'Ping' action will ellicit a 'Pong' response. 
//      Used to keep the manager connection open.
//
func Ping(socket *Socket, actionID string) (bool, error) {

	// verify parameters
	if len(actionID) == 0 {
		return false, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Ping",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	message, err := getMessage(socket, command, actionID)
	if err != nil {
		return false, err
	}
	if message["Response"] != "Success" {
		return false, errors.New(message["Message"])
	}
	return true, nil
}

//	Challenge
//		Generate a challenge for MD5 authentication.
//
func Challenge(socket *Socket, actionID string) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: Challenge",
		"\r\nActionID: ",
		actionID,
		"\r\nAuthType: ",
		"MD5",
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Command
//		Execute Asterisk CLI Command.
//
func Command(socket *Socket, actionID, cmd string) (map[string]string, error) {
	if len(actionID) == 0 || len(cmd) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: Command",
		"\r\nActionID: ",
		actionID,
		"\r\nCommand: ",
		cmd,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	CoreSettings
//		Show PBX core settings (version etc).
//
func CoreSettings(socket *Socket, actionID string) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: CoreSettings",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	CoreStatus
//		Show PBX core status variables.
//
func CoreStatus(socket *Socket, actionID string) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: CoreStatus",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	CreateConfig
//		Creates an empty file in the configuration directory.
//		This action will create an empty file in the configuration directory. 
//		This action is intended to be used before an UpdateConfig action.
//
func CreateConfig(socket *Socket, actionID, filename string) (map[string]string, error) {
	if len(actionID) == 0 || len(filename) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: CreateConfig",
		"\r\nActionID: ",
		actionID,
		"\r\nFilename: ",
		filename,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	DataGet
//		Retrieve the data api tree.
//
func DataGet(socket *Socket, actionID, path, search, filter string) (map[string]string, error) {
	if len(actionID) == 0 || len(path) == 0 ||
		len(search) == 0 || len(filter) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: DataGet",
		"\r\nActionID: ",
		actionID,
		"\r\nPath: ",
		path,
		"\r\nSearch: ",
		search,
		"\r\nFilter: ",
		filter,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Events
//		Control Event Flow.
//		Enable/Disable sending of events to this manager client.
//
func Events(socket *Socket, actionID, eventMask string) (map[string]string, error) {
	if len(actionID) == 0 || len(eventMask) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: Events",
		"\r\nActionID: ",
		actionID,
		"\r\nEventMask: ",
		eventMask,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	GetConfig
//		Retrieve configuration.
//		This action will dump the contents of a configuration file by category and contents or optionally by specified category only.
//
func GetConfig(socket *Socket, actionID, filename, category string) (map[string]string, error) {
	if len(actionID) == 0 || len(filename) == 0 || len(category) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: GetConfig",
		"\r\nActionID: ",
		actionID,
		"\r\nFilename: ",
		filename,
		"\r\nCategory: ",
		category,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	GetConfigJSON
//		Retrieve configuration (JSON format).
//		This action will dump the contents of a configuration file by category and contents in JSON format. 
//		This only makes sense to be used using rawman over the HTTP interface.
//
func GetConfigJSON(socket *Socket, actionID, filename string) (map[string]string, error) {
	if len(actionID) == 0 || len(filename) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: GetConfigJSON",
		"\r\nActionID: ",
		actionID,
		"\r\nFilename: ",
		filename,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	JabberSend
//		Sends a message to a Jabber Client.
//
func JabberSend(socket *Socket, actionID, jabber, jid, message string) (map[string]string, error) {
	if len(actionID) == 0 || len(jabber) == 0 ||
		len(jid) == 0 || len(message) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: JabberSend",
		"\r\nActionID: ",
		actionID,
		"\r\nJabber: ",
		jabber,
		"\r\nJID: ",
		jid,
		"\r\nMessage: ",
		message,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	ListCommands
//		List available manager commands.
//		Returns the action name and synopsis for every action that is available to the user
//
func ListCommands(socket *Socket, actionID string) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: ListCommands",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	ListCategories
//		List categories in configuration file.
//
func ListCategories(socket *Socket, actionID, filename string) (map[string]string, error) {
	if len(actionID) == 0 || len(filename) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: ListCategories",
		"\r\nActionID: ",
		actionID,
		"\r\nFilename: ",
		filename,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
