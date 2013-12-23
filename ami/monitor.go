package ami

import (
	"errors"
)

//	Monitor
//		Monitor a channel.
//		This action may be used to record the audio on a specified channel.
//
func Monitor(socket *Socket, actionID, channel, file, format string, mix bool) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 ||
		len(file) == 0 || len(format) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	m := map[bool]string{false: "false", true: "true"}
	command := []string{
		"Action: Monitor",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nFile: ",
		file,
		"\r\nFormat: ",
		format,
		"\r\nMix: ",
		m[mix],
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	ChangeMonitor
//		Change monitoring filename of a channel.
//		This action may be used to change the file started by a previous 'Monitor' action.
//
func ChangeMonitor(socket *Socket, actionID, channel, file string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 || len(file) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: ChangeMonitor",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nFile: ",
		file,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	MixMonitorMute
//		Mute / unMute a Mixmonitor recording.
//		This action may be used to mute a MixMonitor recording.
//
func MixMonitorMute(socket *Socket, actionID, channel, direction string, state bool) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 || len(direction) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	s := map[bool]string{false: "0", true: "1"}
	command := []string{
		"Action: MixMonitorMute",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nDirection: ",
		direction,
		"\r\nState: ",
		s[state],
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	PauseMonitor
//		Pause monitoring of a channel.
//		This action may be used to temporarily stop the recording of a channel.
//
func PauseMonitor(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: PauseMonitor",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	UnpauseMonitor
//		Unpause monitoring of a channel.
//		This action may be used to re-enable recording of a channel after calling PauseMonitor.
//
func UnpauseMonitor(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: UnpauseMonitor",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	StopMonitor
//		Stop monitoring a channel.
//		This action may be used to end a previously started 'Monitor' action.
//
func StopMonitor(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: StopMonitor",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
