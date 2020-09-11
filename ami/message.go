package ami

import "context"

// MessageSend send an out of call message to an endpoint.
func MessageSend(ctx context.Context, client Client, actionID string, message MessageData) (Response, error) {
	return send(ctx, client, "MessageSend", actionID, message)
}
