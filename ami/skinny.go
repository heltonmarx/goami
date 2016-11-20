package ami

// SKINNYdevices lists SKINNY devices (text format).
// Lists Skinny devices in text format with details on current status.
// Devicelist will follow as separate events,
// followed by a final event called DevicelistComplete.
func SKINNYdevices(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SKINNYdevices",
		"ActionID": actionID,
	})
}

// SKINNYlines lists SKINNY lines (text format).
// Lists Skinny lines in text format with details on current status.
// Linelist will follow as separate events,
// followed by a final event called LinelistComplete.
func SKINNYlines(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SKINNYlines",
		"ActionID": actionID,
	})
}

// SKINNYshowdevice show SKINNY device (text format).
// Show one SKINNY device with details on current status.
func SKINNYshowdevice(socket *Socket, actionID, device string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SKINNYshowdevice",
		"ActionID": actionID,
		"Device":   device,
	})
}

// SKINNYshowline shows SKINNY line (text format).
// Show one SKINNY line with details on current status.
func SKINNYshowline(socket *Socket, actionID, line string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SKINNYshowline",
		"ActionID": actionID,
		"Line":     line,
	})
}
