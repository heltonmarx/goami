package ami

import "strconv"

// KSendSMS sends a SMS using KHOMP device.
func KSendSMS(socket *Socket, actionID string, data KhompSMSData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":       "KSendSMS",
		"ActionID":     actionID,
		"Device":       data.Device,
		"Destination":  data.Destination,
		"Confirmation": strconv.FormatBool(data.Confirmation),
		"Message":      data.Message,
	})
}
