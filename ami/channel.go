package ami

// AbsoluteTimeout set absolute timeout.
// Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
func AbsoluteTimeout(client Client, actionID, channel string, timeout int) (Response, error) {
	return send(client, "AbsoluteTimeout", actionID, map[string]interface{}{
		"Channel": channel,
		"Timeout": timeout,
	})
}

// Atxfer attended transfer.
func Atxfer(client Client, actionID, channel, exten, context, priority string) (Response, error) {
	return send(client, "AbsoluteTimeout", actionID, map[string]string{
		"Channel":  channel,
		"Exten":    exten,
		"Context":  context,
		"Priority": priority,
	})
}

// Bridge bridges two channels already in the PBX.
func Bridge(client Client, actionID, channel1, channel2 string, tone bool) (Response, error) {
	t := map[bool]string{false: "no", true: "yes"}
	return send(client, "Bridge", actionID, map[string]string{
		"Channel1": channel1,
		"Channel2": channel2,
		"Tone":     t[tone],
	})
}

// CoreShowChannels list currently active channels.
func CoreShowChannels(client Client, actionID string) ([]Response, error) {
	return requestList(client, "CoreShowChannels", actionID, "CoreShowChannel", "CoreShowChannelsComplete")
}

// ExtensionState checks extension status.
func ExtensionState(client Client, actionID, exten, context string) (Response, error) {
	return send(client, "ExtensionState", actionID, map[string]string{
		"Exten":   exten,
		"Context": context,
	})
}

// Hangup hangups channel.
func Hangup(client Client, actionID, channel, cause string) (Response, error) {
	return send(client, "Hangup", actionID, map[string]string{
		"Channel": channel,
		"Cause":   cause,
	})
}

// Originate originates a call.
// Generates an outgoing call to a Extension/Context/Priority or Application/Data.
func Originate(client Client, actionID string, originate OriginateData) (Response, error) {
	return send(client, "Originate", actionID, originate)
}

// Park parks a channel.
func Park(client Client, actionID, channel1, channel2 string, timeout int, parkinglot string) (Response, error) {
	return send(client, "Park", actionID, map[string]interface{}{
		"Channel":    channel1,
		"Channel2":   channel2,
		"Timeout":    timeout,
		"Parkinglot": parkinglot,
	})
}

// ParkedCalls list parked calls.
func ParkedCalls(client Client, actionID string) ([]Response, error) {
	return requestList(client, "ParkedCalls", actionID, "ParkedCall", "ParkedCallsComplete")
}

// PlayDTMF plays DTMF signal on a specific channel.
func PlayDTMF(client Client, actionID, channel, digit string) (Response, error) {
	return send(client, "PlayDTMF", actionID, map[string]string{
		"Channel": channel,
		"Digit":   digit,
	})
}

// Redirect redirects (transfer) a call.
func Redirect(client Client, actionID, channel, exten, context, priority string) (Response, error) {
	return send(client, "Redirect", actionID, map[string]string{
		"Channel":  channel,
		"Exten":    exten,
		"Context":  context,
		"Priority": priority,
	})
}

// SendText sends text message to channel.
func SendText(client Client, actionID, channel, msg string) (Response, error) {
	return send(client, "SendText", actionID, map[string]string{
		"Channel": channel,
		"Message": msg,
	})
}

// Setvar sets a channel variable. Sets a global or local channel variable.
// Note: If a channel name is not provided then the variable is global.
func Setvar(client Client, actionID, channel, variable, value string) (Response, error) {
	return send(client, "Setvar", actionID, map[string]string{
		"Channel":  channel,
		"Variable": variable,
		"Value":    value,
	})
}

// Status lists channel status.
// Will return the status information of each channel along with the value for the specified channel variables.
func Status(client Client, actionID, channel, variables string) (Response, error) {
	return send(client, "Status", actionID, map[string]string{
		"Channel":   channel,
		"Variables": variables,
	})
}

// AGI add an AGI command to execute by Async AGI.
func AGI(client Client, actionID, channel, agiCommand, agiCommandID string) (Response, error) {
	return send(client, "AGI", actionID, map[string]string{
		"Channel":     channel,
		"Command: ":   agiCommand,
		"CommandID: ": agiCommandID,
	})
}

// AOCMessage generates an Advice of Charge message on a channel.
func AOCMessage(client Client, actionID string, aocData AOCData) (Response, error) {
	return send(client, "AOCMessage", actionID, aocData)
}

// Getvar gets a channel variable.
func Getvar(client Client, actionID, channel, variable string) (Response, error) {
	return send(client, "Getvar", actionID, map[string]string{
		"Channel":  channel,
		"Variable": variable,
	})
}

// LocalOptimizeAway optimize away a local channel when possible.
// A local channel created with "/n" will not automatically optimize away.
// Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
func LocalOptimizeAway(client Client, actionID, channel string) (Response, error) {
	return send(client, "LocalOptimizeAway", actionID, map[string]string{
		"Channel": channel,
	})
}
