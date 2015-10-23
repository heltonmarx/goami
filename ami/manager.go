// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	errManagerConnectionFailed  = errors.New("manager: AMI connection failed")
	errInvalidManagerParameters = errors.New("manager: Invalid parameters")
)

func VersionInfo() string {
	return "Version: 1.0.0; Build Date: Jan 16, 2014;"
}

func sendCmd(socket *Socket, cmd []string) error {
	for _, s := range cmd {
		if err := socket.Send("%s", s); err != nil {
			fmt.Printf("send login error:[%v]\n", err)
			return err
		}
	}
	return nil
}

func Connect(socket *Socket) (bool, error) {
	if answer, err := socket.Recv(); err != nil || !strings.Contains(answer, "Asterisk Call Manager") {
		return false, errManagerConnectionFailed
	}
	return true, nil
}

//  Login
//      Login Manager.
//
func Login(socket *Socket, user, secret, events, actionID string) (bool, error) {
	// verify parameters
	if len(user) == 0 || len(secret) == 0 || len(events) == 0 || len(actionID) == 0 {
		return false, errInvalidManagerParameters
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
		return false, errInvalidManagerParameters
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
		return false, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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
		return nil, errInvalidManagerParameters
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

//	ModuleCheck
//		Check if module is loaded.
//		Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
//
func ModuleCheck(socket *Socket, actionID, module string) (map[string]string, error) {
	if len(actionID) == 0 || len(module) == 0 {
		return nil, errInvalidManagerParameters
	}
	command := []string{
		"Action: ModuleCheck",
		"\r\nActionID: ",
		actionID,
		"\r\nModule: ",
		module,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	ModuleLoad
//		Module management.
//		Loads, unloads or reloads an Asterisk module in a running system.
//
func ModuleLoad(socket *Socket, actionID, module, loadType string) (map[string]string, error) {
	if len(actionID) == 0 || len(module) == 0 || len(loadType) == 0 {
		return nil, errInvalidManagerParameters
	}
	command := []string{
		"Action: ModuleLoad",
		"\r\nActionID: ",
		actionID,
		"\r\nModule: ",
		module,
		"\r\nLoadType: ",
		loadType,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Reload
//		Send a reload event.
//
func Reload(socket *Socket, actionID, module string) (map[string]string, error) {
	if len(actionID) == 0 || len(module) == 0 {
		return nil, errInvalidManagerParameters
	}
	command := []string{
		"Action: Reload",
		"\r\nActionID: ",
		actionID,
		"\r\nModule: ",
		module,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	ShowDialPlan
//		Show dialplan contexts and extensions
//		Be aware that showing the full dialplan may take a lot of capacity
//
func ShowDialPlan(socket *Socket, actionID, extension, context string) (map[string]string, error) {
	if len(actionID) == 0 || len(extension) == 0 || len(context) == 0 {
		return nil, errInvalidManagerParameters
	}
	command := []string{
		"Action: ShowDialPlan",
		"\r\nActionID: ",
		actionID,
		"\r\nExtension: ",
		extension,
		"\r\nContext: ",
		context,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	GetEvents
//		Get events from current socket connection
//		It is mandatory set 'events' of ami.Login with "system,call,all,user", to received
//		events
//
func GetEvents(socket *Socket) (map[string]string, error) {
	message, err := decode(socket)
	if err != nil {
		return nil, err
	}
	return message, nil
}
