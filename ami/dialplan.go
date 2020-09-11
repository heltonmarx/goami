package ami

import "context"

// DialplanExtensionAdd add an extension to the dialplan.
func DialplanExtensionAdd(ctx context.Context, client Client, actionID string, extension ExtensionData) (Response, error) {
	return send(ctx, client, "DialplanExtensionAdd", actionID, extension)
}

// DialplanExtensionRemove remove an extension from the dialplan.
func DialplanExtensionRemove(ctx context.Context, client Client, actionID string, extension ExtensionData) (Response, error) {
	return send(ctx, client, "DialplanExtensionRemove", actionID, extension)
}
