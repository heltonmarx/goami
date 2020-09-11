package ami

import "context"

// AGIControl represents the control type to playback actions
type AGIControl string

const (
	// Stop the playback operation
	Stop AGIControl = "stop"
	// Forward move the current position in the media forward.
	Forward = "forward"
	// Reverse move the current poistion in the media backward.
	Reverse = "reverse"
	// Pause pause/unpause the playback operation.
	Pause = "pause"
	// Restart the playback operation.
	Restart = "restart"
)

// AGI add an AGI command to execute by Async AGI.
func AGI(ctx context.Context, client Client, actionID, channel, agiCommand, agiCommandID string) (Response, error) {
	return send(ctx, client, "AGI", actionID, map[string]string{
		"Channel":   channel,
		"Command":   agiCommand,
		"CommandID": agiCommandID,
	})
}

// ControlPlayback control the playback of a file being played to a channel.
func ControlPlayback(ctx context.Context, client Client, actionID, channel string, control AGIControl) (Response, error) {
	return send(ctx, client, "ControlPlayback", actionID, map[string]interface{}{
		"Channel": channel,
		"Control": control,
	})
}
