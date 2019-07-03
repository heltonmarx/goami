package ami

// ExtensionState checks extension status.
func ExtensionState(client Client, actionID, exten, context string) (Response, error) {
	return send(client, "ExtensionState", actionID, map[string]string{
		"Exten":   exten,
		"Context": context,
	})
}

// ExtensionStateList list the current known extension states.
func ExtensionStateList(client Client, actionID string) ([]Response, error) {
	return requestList(client, "ExtensionStateList", actionID, "ExtensionStatus", "ExtensionStateListComplete")
}
