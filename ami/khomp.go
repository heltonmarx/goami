package ami

// KSendSMS sends a SMS using KHOMP device.
func KSendSMS(client Client, actionID string, data KhompSMSData) (Response, error) {
	return send(client, "KSendSMS", actionID, data)
}
