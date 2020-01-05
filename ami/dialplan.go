package ami

// DialplanExtensionAdd add an extension to the dialplan.
func DialplanExtensionAdd(client Client, actionID string, extension ExtensionData) (Response, error) {
	return send(client, "DialplanExtensionAdd", actionID, extension)
}

// DialplanExtensionRemove remove an extension from the dialplan.
func DialplanExtensionRemove(client Client, actionID string, extension ExtensionData) (Response, error) {
	return send(client, "DialplanExtensionRemove", actionID, extension)
}
