package ami

import "context"

// Atxfer attended transfer.
func Atxfer(ctx context.Context, client Client, actionID, channel, exten, context string) (Response, error) {
	return send(ctx, client, "Atxfer", actionID, map[string]string{
		"Channel": channel,
		"Exten":   exten,
		"Context": context,
	})
}

// CancelAtxfer cancel an attended transfer.
func CancelAtxfer(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "CancelAtxfer", actionID, map[string]string{
		"Channel": channel,
	})
}
