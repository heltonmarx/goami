package ami

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Connect verify connect response.
func Connect(socket *Socket) (bool, error) {
	if answer, err := socket.Recv(); err != nil || !strings.Contains(answer, "Asterisk Call Manager") {
		return false, errors.New("manager: Invalid parameters")
	}
	return true, nil
}

// Login provides the login manager.
func Login(socket *Socket, user, secret, events, actionID string) error {
	resp, err := sendCommand(socket, map[string]string{
		"Action":   "Login",
		"Username": user,
		"Secret":   secret,
		"Events":   events,
		"ActionID": actionID,
	})
	if err != nil {
		return err
	}
	if msg, ok := resp["Response"]; !ok || msg != "Success" {
		if r, ok := resp["Message"]; ok {
			return errors.New(r)
		}
		return errors.New("login failed")
	}
	return nil
}

// Logoff logoff the current manager session.
func Logoff(socket *Socket, actionID string) error {
	resp, err := sendCommand(socket, map[string]string{
		"Action":   "Logoff",
		"ActionID": actionID,
	})
	if err != nil {
		return err
	}
	if msg, ok := resp["Response"]; !ok || msg != "Goodbye" {
		if r, ok := resp["Message"]; ok {
			return errors.New(r)
		}
		return errors.New("logout failed")
	}
	return nil
}

// GetUUID returns a new UUID based on /dev/urandom (unix).
func GetUUID() (string, error) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		return "", fmt.Errorf("open /dev/urandom error:[%v]", err)
	}
	defer f.Close()
	b := make([]byte, 16)

	_, err = f.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

// Ping action will ellicit a 'Pong' response.
// Used to keep the manager connection open.
func Ping(socket *Socket, actionID string) error {
	resp, err := sendCommand(socket, map[string]string{
		"Action":   "Ping",
		"ActionID": actionID,
	})
	if err != nil {
		return err
	}
	if msg, ok := resp["Response"]; !ok || msg != "Success" {
		if r, ok := resp["Message"]; ok {
			return errors.New(r)
		}
		return errors.New("login failed")
	}
	return nil
}

// Challenge generates a challenge for MD5 authentication.
func Challenge(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Challenge",
		"ActionID": actionID,
		"AuthType": "MD5",
	})
}

// Command executes an Asterisk CLI Command.
func Command(socket *Socket, actionID, cmd string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Command",
		"ActionID": actionID,
		"Command":  cmd,
	})
}

// CoreSettings shows PBX core settings (version etc).
func CoreSettings(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "CoreSettings",
		"ActionID": actionID,
	})
}

// CoreStatus shows PBX core status variables.
func CoreStatus(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "CoreStatus",
		"ActionID": actionID,
	})
}

// CreateConfig creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory.
// This action is intended to be used before an UpdateConfig action.
func CreateConfig(socket *Socket, actionID, filename string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "CreateConfig",
		"ActionID": actionID,
		"Filename": filename,
	})
}

// DataGet retrieves the data api tree.
func DataGet(socket *Socket, actionID, path, search, filter string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "DataGet",
		"ActionID": actionID,
		"Path":     path,
		"Search":   search,
		"Filter":   filter,
	})
}

// Events control Event Flow.
// Enable/Disable sending of events to this manager client.
func Events(socket *Socket, actionID, eventMask string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":    "Events",
		"ActionID":  actionID,
		"EventMask": eventMask,
	})
}

// GetConfig retrieves configuration.
// This action will dump the contents of a configuration file by category and contents or optionally by specified category only.
func GetConfig(socket *Socket, actionID, filename, category string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "GetConfig",
		"ActionID": actionID,
		"Filename": filename,
		"Category": category,
	})
}

// GetConfigJSON retrieves configuration (JSON format).
// This action will dump the contents of a configuration file by category and contents in JSON format.
// This only makes sense to be used using rawman over the HTTP interface.
func GetConfigJSON(socket *Socket, actionID, filename string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "GetConfigJSON",
		"ActionID": actionID,
		"Filename": filename,
	})
}

// JabberSend sends a message to a Jabber Client.
func JabberSend(socket *Socket, actionID, jabber, jid, message string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "JabberSend",
		"ActionID": actionID,
		"Jabber":   jabber,
		"JID":      jid,
		"Message":  message,
	})
}

// ListCommands lists available manager commands.
// Returns the action name and synopsis for every action that is available to the user
func ListCommands(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "ListCommands",
		"ActionID": actionID,
	})
}

// ListCategories lists categories in configuration file.
func ListCategories(socket *Socket, actionID, filename string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "ListCategories",
		"ActionID": actionID,
		"Filename": filename,
	})
}

// ModuleCheck checks if module is loaded.
// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
func ModuleCheck(socket *Socket, actionID, module string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "ModuleCheck",
		"ActionID": actionID,
		"Module":   module,
	})
}

// ModuleLoad module management.
// Loads, unloads or reloads an Asterisk module in a running system.
func ModuleLoad(socket *Socket, actionID, module, loadType string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "ModuleLoad",
		"ActionID": actionID,
		"Module":   module,
		"LoadType": loadType,
	})
}

// Reload Sends a reload event.
func Reload(socket *Socket, actionID, module string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Reload",
		"ActionID": actionID,
		"Module":   module,
	})
}

// ShowDialPlan shows dialplan contexts and extensions
// Be aware that showing the full dialplan may take a lot of capacity.
func ShowDialPlan(socket *Socket, actionID, extension, context string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":    "ShowDialPlan",
		"ActionID":  actionID,
		"Extension": extension,
		"Context":   context,
	})
}

// GetEvents gets events from current socket connection
// It is mandatory set 'events' of ami.Login with "system,call,all,user", to received events.
func GetEvents(socket *Socket) (map[string]string, error) {
	return decode(socket)
}
