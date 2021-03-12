package ami

import (
	"context"
	"strconv"
)

// AbsoluteTimeout set absolute timeout.
// Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
func AbsoluteTimeout(ctx context.Context, client Client, actionID, channel string, timeout int) (Response, error) {
	return send(ctx, client, "AbsoluteTimeout", actionID, map[string]interface{}{
		"Channel": channel,
		"Timeout": timeout,
	})
}

// CoreShowChannels list currently active channels.
func CoreShowChannels(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "CoreShowChannels", actionID, "CoreShowChannel", "CoreShowChannelsComplete")
}

// Hangup hangups channel.
func Hangup(ctx context.Context, client Client, actionID, channel, cause string) (Response, error) {
	return send(ctx, client, "Hangup", actionID, map[string]string{
		"Channel": channel,
		"Cause":   cause,
	})
}

// Originate originates a call.
// Generates an outgoing call to a Extension/Context/Priority or Application/Data.
func Originate(ctx context.Context, client Client, actionID string, originate OriginateData) (Response, error) {
	return send(ctx, client, "Originate", actionID, originate)
}

// Park parks a channel.
func Park(ctx context.Context, client Client, actionID, channel1, channel2 string, timeout int, parkinglot string) (Response, error) {
	return send(ctx, client, "Park", actionID, map[string]interface{}{
		"Channel":    channel1,
		"Channel2":   channel2,
		"Timeout":    timeout,
		"Parkinglot": parkinglot,
	})
}

// ParkedCalls list parked calls.
func ParkedCalls(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "ParkedCalls", actionID, "ParkedCall", "ParkedCallsComplete")
}

// Parkinglots get a list of parking lots.
func Parkinglots(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "Parkinglots", actionID, "ParkedCall", "ParkinglotsComplete")
}

// PlayDTMF plays DTMF signal on a specific channel.
func PlayDTMF(ctx context.Context, client Client, actionID, channel, digit string, duration int) (Response, error) {
	return send(ctx, client, "PlayDTMF", actionID, map[string]string{
		"Channel":  channel,
		"Digit":    digit,
		"Duration": strconv.Itoa(duration),
	})
}

// Redirect redirects (transfer) a call.
func Redirect(ctx context.Context, client Client, actionID string, call CallData) (Response, error) {
	return send(ctx, client, "Redirect", actionID, call)
}

// SendText sends text message to channel.
func SendText(ctx context.Context, client Client, actionID, channel, msg string) (Response, error) {
	return send(ctx, client, "SendText", actionID, map[string]string{
		"Channel": channel,
		"Message": msg,
	})
}

// Setvar sets a channel variable. Sets a global or local channel variable.
// Note: If a channel name is not provided then the variable is global.
func Setvar(ctx context.Context, client Client, actionID, channel, variable, value string) (Response, error) {
	return send(ctx, client, "Setvar", actionID, map[string]string{
		"Channel":  channel,
		"Variable": variable,
		"Value":    value,
	})
}

// Status lists channel status.
// Will return the status information of each channel along with the value for the specified channel variables.
func Status(ctx context.Context, client Client, actionID, channel, variables string) (Response, error) {
	return send(ctx, client, "Status", actionID, map[string]string{
		"Channel":   channel,
		"Variables": variables,
	})
}

// AOCMessage generates an Advice of Charge message on a channel.
func AOCMessage(ctx context.Context, client Client, actionID string, aocData AOCData) (Response, error) {
	return send(ctx, client, "AOCMessage", actionID, aocData)
}

// Getvar gets a channel variable.
func Getvar(ctx context.Context, client Client, actionID, channel, variable string) (Response, error) {
	return send(ctx, client, "Getvar", actionID, map[string]string{
		"Channel":  channel,
		"Variable": variable,
	})
}

// LocalOptimizeAway optimize away a local channel when possible.
// A local channel created with "/n" will not automatically optimize away.
// Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
func LocalOptimizeAway(ctx context.Context, client Client, actionID, channel string) (Response, error) {
	return send(ctx, client, "LocalOptimizeAway", actionID, map[string]string{
		"Channel": channel,
	})
}

// MuteAudio mute an audio stream.
func MuteAudio(ctx context.Context, client Client, actionID, channel, direction string, state bool) (Response, error) {
	stateMap := map[bool]string{false: "off", true: "on"}
	return send(ctx, client, "MuteAudio", actionID, map[string]string{
		"Channel":   channel,
		"Direction": direction,
		"State":     stateMap[state],
	})
}
