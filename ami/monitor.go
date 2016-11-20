package ami

import "strconv"

// Monitor monitors a channel.
// This action may be used to record the audio on a specified channel.
func Monitor(socket *Socket, actionID, channel, file, format string, mix bool) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Monitor",
		"ActionID": actionID,
		"Channel":  channel,
		"File":     file,
		"Format":   format,
		"Mix":      strconv.FormatBool(mix),
	})
}

// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func ChangeMonitor(socket *Socket, actionID, channel, file string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "ChangeMonitor",
		"ActionID": actionID,
		"Channel":  channel,
		"File: ":   file,
	})
}

// MixMonitorMute Mute / unMute a Mixmonitor recording.
// This action may be used to mute a MixMonitor recording.
func MixMonitorMute(socket *Socket, actionID, channel, direction string, state bool) (map[string]string, error) {
	s := map[bool]string{false: "0", true: "1"}
	return sendCommand(socket, map[string]string{
		"Action":    "MixMonitorMute",
		"ActionID":  actionID,
		"Channel":   channel,
		"Direction": direction,
		"State":     s[state],
	})
}

// PauseMonitor pauses monitoring of a channel.
// This action may be used to temporarily stop the recording of a channel.
func PauseMonitor(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "PauseMonitor",
		"ActionID": actionID,
		"Channel":  channel,
	})
}

// UnpauseMonitor unpause monitoring of a channel.
// This action may be used to re-enable recording of a channel after calling PauseMonitor.
func UnpauseMonitor(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "UnpauseMonitor",
		"ActionID": actionID,
		"Channel":  channel,
	})
}

// StopMonitor stops monitoring a channel.
// This action may be used to end a previously started 'Monitor' action.
func StopMonitor(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "StopMonitor",
		"ActionID": actionID,
		"Channel":  channel,
	})
}
