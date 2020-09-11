package ami

import "context"

// ExtensionState checks extension status.
func ExtensionState(ctx context.Context, client Client, actionID, exten, context string) (Response, error) {
	return send(ctx, client, "ExtensionState", actionID, map[string]string{
		"Exten":   exten,
		"Context": context,
	})
}

// ExtensionStateList list the current known extension states.
func ExtensionStateList(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "ExtensionStateList", actionID, "ExtensionStatus", "ExtensionStateListComplete")
}
