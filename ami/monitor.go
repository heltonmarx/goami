package ami

// Monitor monitors a channel.
// This action may be used to record the audio on a specified channel.
func Monitor(client Client, actionID, channel, file, format string, mix bool) (Response, error) {
	return send(client, "Monitor", actionID, monitorData{
		Channel: channel,
		File:    file,
		Format:  format,
		Mix:     mix,
	})
}

// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func ChangeMonitor(client Client, actionID, channel, file string) (Response, error) {
	return send(client, "ChangeMonitor", actionID, monitorData{
		Channel: channel,
		File:    file,
	})
}

// MixMonitorMute Mute / unMute a Mixmonitor recording.
// This action may be used to mute a MixMonitor recording.
func MixMonitorMute(client Client, actionID, channel, direction string, state bool) (Response, error) {
	s := map[bool]string{false: "0", true: "1"}
	return send(client, "MixMonitorMute", actionID, monitorData{
		Channel:   channel,
		Direction: direction,
		State:     s[state],
	})
}

// PauseMonitor pauses monitoring of a channel.
// This action may be used to temporarily stop the recording of a channel.
func PauseMonitor(client Client, actionID, channel string) (Response, error) {
	return send(client, "PauseMonitor", actionID, monitorData{
		Channel: channel,
	})
}

// UnpauseMonitor unpause monitoring of a channel.
// This action may be used to re-enable recording of a channel after calling PauseMonitor.
func UnpauseMonitor(client Client, actionID, channel string) (Response, error) {
	return send(client, "UnpauseMonitor", actionID, monitorData{
		Channel: channel,
	})
}

// StopMonitor stops monitoring a channel.
// This action may be used to end a previously started 'Monitor' action.
func StopMonitor(client Client, actionID, channel string) (Response, error) {
	return send(client, "StopMonitor", actionID, monitorData{
		Channel: channel,
	})
}

type monitorData struct {
	Channel   string `ami:"Channel"`
	Direction string `ami:"Direction,omitempty"`
	State     string `ami:"State,omitempty"`
	File      string `ami:"File, omitempty"`
	Format    string `ami:"Format,omitempty"`
	Mix       bool   `ami:"Mix,omitempty"`
}
