package ami

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// Connect verify connect response.
func Connect(ctx context.Context, client Client) (bool, error) {
	if answer, err := client.Recv(ctx); err != nil || !strings.Contains(answer, "Asterisk Call Manager") {
		return false, errors.New("manager: Invalid parameters")
	}
	return true, nil
}

// Login provides the login manager.
func Login(ctx context.Context, client Client, user, secret, events, actionID string) error {
	var login = struct {
		Username string `ami:"Username"`
		Secret   string `ami:"Secret"`
		Events   string `ami:"Events,omitempty"`
	}{Username: user, Secret: secret, Events: events}
	resp, err := send(ctx, client, "Login", actionID, login)
	if err != nil {
		return err
	}
	if ok := resp.Get("Response"); ok != "Success" {
		return fmt.Errorf("login failed: %v", resp.Get("Message"))
	}
	return nil
}

// Logoff logoff the current manager session.
func Logoff(ctx context.Context, client Client, actionID string) error {
	resp, err := send(ctx, client, "Logoff", actionID, nil)
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
func Ping(ctx context.Context, client Client, actionID string) error {
	resp, err := send(ctx, client, "Ping", actionID, nil)
	if err != nil {
		return err
	}
	if ok := resp.Get("Response"); ok != "Success" {
		return fmt.Errorf("ping failed: %v", resp.Get("Message"))
	}
	return nil
}

// Challenge generates a challenge for MD5 authentication.
func Challenge(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "Challenge", actionID, map[string]string{
		"AuthType": "MD5",
	})
}

// Command executes an Asterisk CLI Command.
func Command(ctx context.Context, client Client, actionID, cmd string) (Response, error) {
	return send(ctx, client, "Command", actionID, map[string]string{
		"Command": cmd,
	})
}

// CoreSettings shows PBX core settings (version etc).
func CoreSettings(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "CoreSettings", actionID, nil)
}

// CoreStatus shows PBX core status variables.
func CoreStatus(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "CoreStatus", actionID, nil)
}

// CreateConfig creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory.
// This action is intended to be used before an UpdateConfig action.
func CreateConfig(ctx context.Context, client Client, actionID, filename string) (Response, error) {
	return send(ctx, client, "CreateConfig", actionID, map[string]string{
		"Filename": filename,
	})
}

// DataGet retrieves the data api tree.
func DataGet(ctx context.Context, client Client, actionID, path, search, filter string) (Response, error) {
	return send(ctx, client, "DataGet", actionID, map[string]string{
		"Path":   path,
		"Search": search,
		"Filter": filter,
	})
}

// EventFlow control Event Flow.
// Enable/Disable sending of events to this manager client.
func EventFlow(ctx context.Context, client Client, actionID, eventMask string) (Response, error) {
	return send(ctx, client, "Events", actionID, map[string]string{
		"EventMask": eventMask,
	})
}

// GetConfig retrieves configuration.
// This action will dump the contents of a configuration file by category and contents or optionally by specified category only.
func GetConfig(ctx context.Context, client Client, actionID, filename, category, filter string) (Response, error) {
	return send(ctx, client, "GetConfig", actionID, map[string]string{
		"Filename": filename,
		"Category": category,
		"Filter":   filter,
	})
}

// GetConfigJSON retrieves configuration (JSON format).
// This action will dump the contents of a configuration file by category and contents in JSON format.
// This only makes sense to be used using rawman over the HTTP interface.
func GetConfigJSON(ctx context.Context, client Client, actionID, filename, category, filter string) (Response, error) {
	return send(ctx, client, "GetConfigJSON", actionID, map[string]string{
		"Filename": filename,
		"Category": category,
		"Filter":   filter,
	})
}

// JabberSend sends a message to a Jabber Client
func JabberSend(ctx context.Context, client Client, actionID, jabber, jid, message string) (Response, error) {
	return send(ctx, client, "JabberSend", actionID, map[string]interface{}{
		"Jabber":  jabber,
		"JID":     jid,
		"Message": message,
	})
}

// ListCommands lists available manager commands.
// Returns the action name and synopsis for every action that is available to the user
func ListCommands(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "ListCommands", actionID, nil)
}

// ListCategories lists categories in configuration file.
func ListCategories(ctx context.Context, client Client, actionID, filename string) (Response, error) {
	return send(ctx, client, "ListCategories", actionID, map[string]string{
		"Filename": filename,
	})
}

// ModuleCheck checks if module is loaded.
// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
func ModuleCheck(ctx context.Context, client Client, actionID, module string) (Response, error) {
	return send(ctx, client, "ModuleCheck", actionID, map[string]string{
		"Module": module,
	})
}

// ModuleLoad module management.
// Loads, unloads or reloads an Asterisk module in a running system.
func ModuleLoad(ctx context.Context, client Client, actionID, module, loadType string) (Response, error) {
	return send(ctx, client, "ModuleLoad", actionID, map[string]string{
		"Module":   module,
		"LoadType": loadType,
	})
}

// Reload Sends a reload event.
func Reload(ctx context.Context, client Client, actionID, module string) (Response, error) {
	return send(ctx, client, "Reload", actionID, map[string]string{
		"Module": module,
	})
}

// ShowDialPlan shows dialplan contexts and extensions
// Be aware that showing the full dialplan may take a lot of capacity.
func ShowDialPlan(ctx context.Context, client Client, actionID, extension, context string) (Response, error) {
	return send(ctx, client, "ShowDialPlan", actionID, map[string]string{
		"Extension": extension,
		"Context":   context,
	})
}

// Events gets events from current client connection
// It is mandatory set 'events' of ami.Login with "system,call,all,user", to received events.
func Events(ctx context.Context, client Client) (Response, error) {
	return read(ctx, client)
}

// Filter dinamically add filters for the current manager session.
func Filter(ctx context.Context, client Client, actionID, operation, filter string) (Response, error) {
	return send(ctx, client, "Filter", actionID, map[string]string{
		"Operation": operation,
		"Filter":    filter,
	})
}

// DeviceStateList list the current known device states.
func DeviceStateList(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "DeviceStateList", actionID,
		"DeviceStateChange", "DeviceStateListComplete")
}

// LoggerRotate reload and rotate the Asterisk logger.
func LoggerRotate(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "LoggerRotate", actionID, nil)
}

// UpdateConfig Updates a config file.
// Dynamically updates an Asterisk configuration file.
func UpdateConfig(ctx context.Context, client Client, actionID, srcFilename, dstFilename string, actions []UpdateConfigAction, reload bool) (Response, error) {
	options := make(map[string]string)
	options["SrcFilename"] = srcFilename
	options["DstFilename"] = dstFilename
	if reload {
		options["Reload"] = "yes"
	}
	for i, a := range actions {
		actionNumber := fmt.Sprintf("%06d", i)
		options["Action-"+actionNumber] = a.Action
		options["Cat-"+actionNumber] = a.Category
		if a.Var != "" {
			options["Var-"+actionNumber] = a.Var
			options["Value-"+actionNumber] = a.Value

		}
	}
	return send(ctx, client, "UpdateConfig", actionID, options)
}
