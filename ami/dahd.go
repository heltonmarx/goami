package ami

// DAHDIDialOffhook dials over DAHDI channel while offhook.
// Generate DTMF control frames to the bridged peer.
func DAHDIDialOffhook(socket *Socket, actionID, channel, number string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":       "DAHDIDialOffhook",
		"ActionID":     actionID,
		"DAHDIChannel": channel,
		"Number":       number,
	})
}

// DAHDIDNDoff toggles DAHDI channel Do Not Disturb status OFF.
func DAHDIDNDoff(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":       "DAHDIDNDoff",
		"ActionID":     actionID,
		"DAHDIChannel": channel,
	})
}

// DAHDIDNDon toggles DAHDI channel Do Not Disturb status ON.
func DAHDIDNDon(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":       "DAHDIDNDon",
		"ActionID":     actionID,
		"DAHDIChannel": channel,
	})
}

// DAHDIHangup hangups DAHDI Channel.
func DAHDIHangup(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":       "DAHDIHangup",
		"ActionID":     actionID,
		"DAHDIChannel": channel,
	})
}

// DAHDIRestart fully Restart DAHDI channels (terminates calls).
func DAHDIRestart(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "DAHDIRestart",
		"ActionID": actionID,
	})
}

// DAHDIShowChannels show status of DAHDI channels.
func DAHDIShowChannels(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":       "DAHDIShowChannels",
		"ActionID":     actionID,
		"DAHDIChannel": channel,
	})
}

// DAHDITransfer transfers DAHDI Channel.
func DAHDITransfer(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":       "DAHDITransfer",
		"ActionID":     actionID,
		"DAHDIChannel": channel,
	})
}
