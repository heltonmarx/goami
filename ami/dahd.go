package ami

// DAHDIDialOffhook dials over DAHDI channel while offhook.
// Generate DTMF control frames to the bridged peer.
func DAHDIDialOffhook(client Client, actionID, channel, number string) (Response, error) {
	return send(client, "DAHDIDialOffhook", actionID, map[string]string{
		"DAHDIChannel": channel,
		"Number":       number,
	})
}

// DAHDIDNDoff toggles DAHDI channel Do Not Disturb status OFF.
func DAHDIDNDoff(client Client, actionID, channel string) (Response, error) {
	return send(client, "DAHDIDNDoff", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDIDNDon toggles DAHDI channel Do Not Disturb status ON.
func DAHDIDNDon(client Client, actionID, channel string) (Response, error) {
	return send(client, "DAHDIDNDon", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDIHangup hangups DAHDI Channel.
func DAHDIHangup(client Client, actionID, channel string) (Response, error) {
	return send(client, "DAHDIHangup", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDIRestart fully Restart DAHDI channels (terminates calls).
func DAHDIRestart(client Client, actionID string) (Response, error) {
	return send(client, "DAHDIRestart", actionID, nil)
}

// DAHDIShowChannels show status of DAHDI channels.
func DAHDIShowChannels(client Client, actionID, channel string) ([]Response, error) {
	return requestList(client, "DAHDIShowChannels", actionID, "DAHDIShowChannels", "DAHDIShowChannelsComplete", map[string]string{
		"DAHDIChannel": channel,
	})
}

// DAHDITransfer transfers DAHDI Channel.
func DAHDITransfer(client Client, actionID, channel string) (Response, error) {
	return send(client, "DAHDITransfer", actionID, map[string]string{
		"DAHDIChannel": channel,
	})
}
