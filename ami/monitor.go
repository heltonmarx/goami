package ami

import "context"

// Monitor monitors a channel.
// This action may be used to record the audio on a specified channel.
func Monitor(ctx context.Context, client Client, actionID, channel, file, format string, mix bool) (Response, error) {
	return send(ctx, client, "Monitor", actionID, monitorData{
		Channel: channel,
		File:    file,
		Format:  format,
		Mix:     mix,
	})
}

// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func ChangeMonitor(ctx context.Context, client Client, actionID, channel, file string) (Response, error) {
	return send(ctx, client, "ChangeMonitor", actionID, monitorData{
		Channel: channel,
		File:    file,
	})
}

// MixMonitor record a call and mix the audio during the recording.
func MixMonitor(ctx context.Context, client Client, actionID, channel, file, options, command string) (Response, error) {
	return send(ctx, client, "MixMonitor", actionID, map[string]string{
		"Channel": channel,
		"File":    file,
		"options": options,
		"Command": command,
	})
}

// MixMonitorMute Mute / unMute a Mixmonitor recording.
// This action may be used to mute a MixMonitor recording.
func MixMonitorMute(ctx context.Context, client Client, actionID, channel, direction string, state bool) (Response, error) {
	s := map[bool]string{false: "0", true: "1"}
	return send(ctx, client, "MixMonitorMute", actionID, monitorData{
		Channel:   channel,
		Direction: direction,
		State:     s[state],
	})
}

// PauseMonitor pauses monitoring of a channel.
// This action may be used to temporarily stop the recording of a channel.
func PauseMonitor(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "PauseMonitor", actionID, monitorData{
		Channel: channel,
	})
}

// UnpauseMonitor unpause monitoring of a channel.
// This action may be used to re-enable recording of a channel after calling PauseMonitor.
func UnpauseMonitor(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "UnpauseMonitor", actionID, monitorData{
		Channel: channel,
	})
}

// StopMonitor stops monitoring a channel.
// This action may be used to end a previously started 'Monitor' action.
func StopMonitor(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "StopMonitor", actionID, monitorData{
		Channel: channel,
	})
}

// StopMixMonitor stop recording a call through MixMonitor, and free the recording's file handle.
func StopMixMonitor(ctx context.Context, client Client, actionID, channel, mixMonitorID string) (Response, error) {
	return send(ctx, client, "StopMixMonitor", actionID, monitorData{
		Channel:      channel,
		MixMonitorID: mixMonitorID,
	})
}

type monitorData struct {
	Channel      string `ami:"Channel"`
	Direction    string `ami:"Direction,omitempty"`
	State        string `ami:"State,omitempty"`
	File         string `ami:"File, omitempty"`
	Format       string `ami:"Format,omitempty"`
	Mix          bool   `ami:"Mix,omitempty"`
	MixMonitorID string `ami:"MixMonitorID,omitempty"`
}
