package ami

import "context"

// DAHDIDialOffhook dials over DAHDI channel while offhook.
// Generate DTMF control frames to the bridged peer.
func DAHDIDialOffhook(ctx context.Context, client Client, actionID, channel, number string) (Response, error) {
	return send(ctx, client, "DAHDIDialOffhook", actionID, map[string]string{
		"DAHDIChannel": channel,
		"Number":       number,
	})
}

// DAHDIDNDoff toggles DAHDI channel Do Not Disturb status OFF.
func DAHDIDNDoff(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "DAHDIDNDoff", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDIDNDon toggles DAHDI channel Do Not Disturb status ON.
func DAHDIDNDon(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "DAHDIDNDon", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDIHangup hangups DAHDI Channel.
func DAHDIHangup(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "DAHDIHangup", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDIRestart fully Restart DAHDI channels (terminates calls).
func DAHDIRestart(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "DAHDIRestart", actionID, nil)
}

// DAHDIShowChannels show status of DAHDI channels.
func DAHDIShowChannels(ctx context.Context, client Client, actionID, channel string) ([]Response, error) {
	return requestList(ctx, client, "DAHDIShowChannels", actionID, "DAHDIShowChannels", "DAHDIShowChannelsComplete", map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDITransfer transfers DAHDI Channel.
func DAHDITransfer(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "DAHDITransfer", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}
