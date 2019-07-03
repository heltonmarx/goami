package ami

import (
	"errors"
	"fmt"
	"strings"
)

// Connect verify connect response.
func Connect(client Client) (bool, error) {
	if answer, err := client.Recv(); err != nil || !strings.Contains(answer, "Asterisk Call Manager") {
		return false, errors.New("manager: Invalid parameters")
	}
	return true, nil
}

// Login provides the login manager.
func Login(client Client, user, secret, events, actionID string) error {
	var login = struct {
		Username string `ami:"Username"`
		Secret   string `ami:"Secret"`
		Events   string `ami:"Events,omitempty"`
	}{Username: user, Secret: secret, Events: events}
	resp, err := send(client, "Login", actionID, login)
	if err != nil {
		return err
	}
	if ok := resp.Get("Response"); ok != "Success" {
		return fmt.Errorf("login failed: %v", resp.Get("Message"))
	}
	return nil
}

// Logoff logoff the current manager session.
func Logoff(client Client, actionID string) error {
	resp, err := send(client, "Logoff", actionID, nil)
	if err != nil {
		return err
	}
	if msg := resp.Get("Response"); msg != "Goodbye" {
		return fmt.Errorf("logout failed: %v", resp.Get("Message"))
	}
	return nil
}

// Ping action will ellicit a 'Pong' response.
// Used to keep the manager connection open.
func Ping(client Client, actionID string) error {
	resp, err := send(client, "Ping", actionID, nil)
	if err != nil {
		return err
	}
	if ok := resp.Get("Response"); ok != "Success" {
		return fmt.Errorf("ping failed: %v", resp.Get("Message"))
	}
	return nil
}

// Challenge generates a challenge for MD5 authentication.
func Challenge(client Client, actionID string) (Response, error) {
	return send(client, "Challenge", actionID, map[string]string{
		"AuthType": "MD5",
	})
}

// Command executes an Asterisk CLI Command.
func Command(client Client, actionID, cmd string) (Response, error) {
	return send(client, "Command", actionID, map[string]string{
		"Command": cmd,
	})
}

// CoreSettings shows PBX core settings (version etc).
func CoreSettings(client Client, actionID string) (Response, error) {
	return send(client, "CoreSettings", actionID, nil)
}

// CoreStatus shows PBX core status variables.
func CoreStatus(client Client, actionID string) (Response, error) {
	return send(client, "CoreStatus", actionID, nil)
}

// CreateConfig creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory.
// This action is intended to be used before an UpdateConfig action.
func CreateConfig(client Client, actionID, filename string) (Response, error) {
	return send(client, "CreateConfig", actionID, map[string]string{
		"Filename": filename,
	})
}

// DataGet retrieves the data api tree.
func DataGet(client Client, actionID, path, search, filter string) (Response, error) {
	return send(client, "DataGet", actionID, map[string]string{
		"Path":   path,
		"Search": search,
		"Filter": filter,
	})
}

// EventFlow control Event Flow.
// Enable/Disable sending of events to this manager client.
func EventFlow(client Client, actionID, eventMask string) (Response, error) {
	return send(client, "Events", actionID, map[string]string{
		"EventMask": eventMask,
	})
}

// GetConfig retrieves configuration.
// This action will dump the contents of a configuration file by category and contents or optionally by specified category only.
func GetConfig(client Client, actionID, filename, category, filter string) (Response, error) {
	return send(client, "GetConfig", actionID, map[string]string{
		"Filename": filename,
		"Category": category,
		"Filter":   filter,
	})
}

// GetConfigJSON retrieves configuration (JSON format).
// This action will dump the contents of a configuration file by category and contents in JSON format.
// This only makes sense to be used using rawman over the HTTP interface.
func GetConfigJSON(client Client, actionID, filename, category, filter string) (Response, error) {
	return send(client, "GetConfigJSON", actionID, map[string]string{
		"Filename": filename,
		"Category": category,
		"Filter":   filter,
	})
}

// JabberSend sends a message to a Jabber Client.
func JabberSend(client Client, actionID, jabber, jid, message string) (Response, error) {
	return send(client, "JabberSend", actionID, map[string]string{
		"Jabber":  jabber,
		"JID":     jid,
		"Message": message,
	})
}

// ListCommands lists available manager commands.
// Returns the action name and synopsis for every action that is available to the user
func ListCommands(client Client, actionID string) (Response, error) {
	return send(client, "ListCommands", actionID, nil)
}

// ListCategories lists categories in configuration file.
func ListCategories(client Client, actionID, filename string) (Response, error) {
	return send(client, "ListCategories", actionID, map[string]string{
		"Filename": filename,
	})
}

// ModuleCheck checks if module is loaded.
// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
func ModuleCheck(client Client, actionID, module string) (Response, error) {
	return send(client, "ModuleCheck", actionID, map[string]string{
		"Module": module,
	})
}

// ModuleLoad module management.
// Loads, unloads or reloads an Asterisk module in a running system.
func ModuleLoad(client Client, actionID, module, loadType string) (Response, error) {
	return send(client, "ModuleLoad", actionID, map[string]string{
		"Module":   module,
		"LoadType": loadType,
	})
}

// Reload Sends a reload event.
func Reload(client Client, actionID, module string) (Response, error) {
	return send(client, "Reload", actionID, map[string]string{
		"Module": module,
	})
}

// ShowDialPlan shows dialplan contexts and extensions
// Be aware that showing the full dialplan may take a lot of capacity.
func ShowDialPlan(client Client, actionID, extension, context string) (Response, error) {
	return send(client, "ShowDialPlan", actionID, map[string]string{
		"Extension": extension,
		"Context":   context,
	})
}

// Events gets events from current client connection
// It is mandatory set 'events' of ami.Login with "system,call,all,user", to received events.
func Events(client Client) (Response, error) {
	return read(client)
}

// Filter dinamically add filters for the current manager session.
func Filter(client Client, actionID, operation, filter string) (Response, error) {
	return send(client, "Filter", actionID, map[string]string{
		"Operation": operation,
		"Filter":    filter,
	})
}

// DeviceStateList list the current known device states.
func DeviceStateList(client Client, actionID string) ([]Response, error) {
	return requestList(client, "DeviceStateList", actionID,
		"DeviceStateChange", "DeviceStateListComplete")
}

// LoggerRotate reload and rotate the Asterisk logger.
func LoggerRotate(client Client, actionID string) (Response, error) {
	return send(client, "LoggerRotate", actionID, nil)
}

// DeviceStateList list the current known device states.
func DeviceStateList(client Client, actionID string) ([]Response, error) {
	return requestList(client, "DeviceStateList", actionID,
		"DeviceStateChange", "DeviceStateListComplete")
}

// LoggerRotate reload and rotate the Asterisk logger.
func LoggerRotate(client Client, actionID string) (Response, error) {
	return send(client, "LoggerRotate", actionID, nil)
}
