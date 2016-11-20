package ami

import "strconv"

// AbsoluteTimeout set absolute timeout.
// Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
func AbsoluteTimeout(socket *Socket, actionID, channel string, timeout int) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "AbsoluteTimeout",
		"ActionID": actionID,
		"Channel":  channel,
		"Timeout":  strconv.Itoa(timeout),
	})
}

// Atxfer attended transfer.
func Atxfer(socket *Socket, actionID, channel, exten, context, priority string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "AbsoluteTimeout",
		"ActionID": actionID,
		"Channel":  channel,
		"Exten":    exten,
		"Context":  context,
		"Priority": priority,
	})
}

// Bridge bridges two channels already in the PBX.
func Bridge(socket *Socket, actionID, channel1, channel2 string, tone bool) (map[string]string, error) {
	t := map[bool]string{false: "no", true: "yes"}
	return sendCommand(socket, map[string]string{
		"Action":   "Bridge",
		"ActionID": actionID,
		"Channel1": channel1,
		"Channel2": channel2,
		"Tone":     t[tone],
	})
}

// CoreShowChannels list currently active channels.
func CoreShowChannels(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "CoreShowChannels",
		"ActionID": actionID,
	})
}

// ExtensionState checks extension status.
func ExtensionState(socket *Socket, actionID, exten, context string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "ExtensionState",
		"ActionID": actionID,
		"Exten":    exten,
		"Context":  context,
	})
}

// Hangup hangups channel.
func Hangup(socket *Socket, actionID, channel, cause string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Hangup",
		"ActionID": actionID,
		"Channel":  channel,
		"Cause":    cause,
	})
}

// Originate originates a call.
// Generates an outgoing call to a Extension/Context/Priority or Application/Data.
func Originate(socket *Socket, actionID string, originate OriginateData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":      "Originate",
		"ActionID":    actionID,
		"Channel":     originate.Channel,
		"Exten":       originate.Exten,
		"Context":     originate.Context,
		"Priority":    strconv.Itoa(originate.Priority),
		"Application": originate.Application,
		"Data":        originate.Data,
		"Timeout":     strconv.Itoa(originate.Timeout),
		"CallerID":    originate.Callerid,
		"Variable":    originate.Variable,
		"Account":     originate.Account,
		"Async":       originate.Async,
		"Codecs":      originate.Codecs,
	})
}

// Park parks a channel.
func Park(socket *Socket, actionID, channel1, channel2 string, timeout int, parkinglot string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":     "Park",
		"ActionID":   actionID,
		"Channel":    channel1,
		"Channel2":   channel2,
		"Timeout":    strconv.Itoa(timeout),
		"Parkinglot": parkinglot,
	})
}

// ParkedCalls list parked calls.
func ParkedCalls(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "ParkedCalls",
		"ActionID": actionID,
	})
}

// PlayDTMF plays DTMF signal on a specific channel.
func PlayDTMF(socket *Socket, actionID, channel, digit string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "PlayDTMF",
		"ActionID": actionID,
		"Channel":  channel,
		"Digit":    digit,
	})
}

// Redirect redirects (transfer) a call.
func Redirect(socket *Socket, actionID, channel, exten, context, priority string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Redirect",
		"ActionID": actionID,
		"Channel":  channel,
		"Exten: ":  exten,
		"Context":  context,
		"Priority": priority,
	})
}

// SendText sends text message to channel.
func SendText(socket *Socket, actionID, channel, msg string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SendText",
		"ActionID": actionID,
		"Channel":  channel,
		"Message":  msg,
	})
}

// Setvar sets a channel variable. Sets a global or local channel variable.
// Note: If a channel name is not provided then the variable is global.
func Setvar(socket *Socket, actionID, channel, variable, value string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Setvar",
		"ActionID": actionID,
		"Channel":  channel,
		"Variable": variable,
		"Value":    value,
	})
}

// Status lists channel status.
// Will return the status information of each channel along with the value for the specified channel variables.
func Status(socket *Socket, actionID, channel, variables string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":    "Status",
		"ActionID":  actionID,
		"Channel":   channel,
		"Variables": variables,
	})
}

// AGI add an AGI command to execute by Async AGI.
func AGI(socket *Socket, actionID, channel, agiCommand, agiCommandID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":      "AGI",
		"ActionID: ":  actionID,
		"Channel":     channel,
		"Command: ":   agiCommand,
		"CommandID: ": agiCommandID,
	})
}

// AOCMessage generates an Advice of Charge message on a channel.
func AOCMessage(socket *Socket, actionID string, aocData AOCData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":                    "AOCMessage",
		"ActionID":                  actionID,
		"Channel":                   aocData.Channel,
		"ChannelPrefix":             aocData.ChannelPrefix,
		"MsgType":                   aocData.MsgType,
		"ChargeType":                aocData.ChargeType,
		"UnitAmount(0)":             aocData.UnitAmount,
		"UnitType(0)":               aocData.UnitType,
		"CurrencyName":              aocData.CurrencyName,
		"CurrencyAmount":            aocData.CurrencyAmount,
		"CurrencyMultiplier":        aocData.CurrencyMultiplier,
		"TotalType":                 aocData.TotalType,
		"AOCBillingId":              aocData.AocBillingID,
		"ChargingAssociationId":     aocData.ChargingAssociationID,
		"ChargingAssociationNumber": aocData.ChargingAssociationNumber,
		"ChargingAssociationPlan":   aocData.ChargingAssociationPlan,
	})
}

// Getvar gets a channel variable.
func Getvar(socket *Socket, actionID, channel, variable string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Getvar",
		"ActionID": actionID,
		"Channel":  channel,
		"Variable": variable,
	})
}

// LocalOptimizeAway optimize away a local channel when possible.
// A local channel created with "/n" will not automatically optimize away.
// Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
func LocalOptimizeAway(socket *Socket, actionID, channel string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action:":  "LocalOptimizeAway",
		"ActionID": actionID,
		"Channel":  channel,
	})
}
