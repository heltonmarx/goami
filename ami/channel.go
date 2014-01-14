package ami

import (
	"errors"
	"strconv"
)

//	AbsoluteTimeout	
//		Set absolute timeout.
//		Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
//
func AbsoluteTimeout(socket *Socket, actionID, channel string, timeout int) (map[string]string, error) {
	// verify channel and action ID
	if len(channel) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: AbsoluteTimeout",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nTimeout: ",
		strconv.Itoa(timeout),
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Atxfer
//		Attended transfer.
//
func Atxfer(socket *Socket, actionID, channel, exten, context, priority string) (map[string]string, error) {
	// verify channel and action ID
	if len(channel) == 0 || len(actionID) == 0 ||
		len(exten) == 0 || len(context) == 0 || len(priority) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: AbsoluteTimeout",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nExten: ",
		exten,
		"\r\nContext: ",
		context,
		"\r\nPriority: ",
		priority,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Bridge
//		Bridge two channels already in the PBX.
//	
func Bridge(socket *Socket, actionID, channel1, channel2 string, tone bool) (map[string]string, error) {
	// verify channel and action ID
	if len(channel1) == 0 || len(channel2) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	t := map[bool]string{false: "no", true: "yes"}

	command := []string{
		"Action: Bridge",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel1: ",
		channel1,
		"\r\nChannel2: ",
		channel2,
		"\r\nTone: ",
		t[tone],
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	CoreShowChannels
//		List currently active channels.
//
func CoreShowChannels(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("CoreShowChannels", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"CoreShowChannel",
		"CoreShowChannelsComplete")
}

//
//	ExtensionState
//		Check Extension Status.
//
func ExtensionState(socket *Socket, actionID, exten, context string) (map[string]string, error) {
	// verify action ID
	if len(actionID) == 0 || len(exten) == 0 || len(context) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: ExtensionState",
		"\r\nActionID: ",
		actionID,
		"\r\nExten: ",
		exten,
		"\r\nContext: ",
		context,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Hangup
//		Hangup channel.
//
func Hangup(socket *Socket, actionID, channel, cause string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Hangup",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nCause: ",
		cause,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Originate
//		Originate a call.
//		Generates an outgoing call to a Extension/Context/Priority or Application/Data
//
func Originate(socket *Socket, actionID string, originate OriginateData) (map[string]string, error) {
	// verify action ID
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Originate",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		originate.Channel,
		"\r\nExten: ",
		originate.Exten,
		"\r\nContext: ",
		originate.Context,
		"\n\rPriority: ",
		strconv.Itoa(originate.Priority),
		"\r\nApplication: ",
		originate.Application,
		"\r\nData: ",
		originate.Data,
		"\r\nTimeout: ",
		strconv.Itoa(originate.Timeout),
		"\r\nCallerID: ",
		originate.Callerid,
		"\r\nVariable: ",
		originate.Variable,
		"\r\nAccount: ",
		originate.Account,
		"\r\nAsync: ",
		originate.Async,
		"\r\nCodecs: ",
		originate.Codecs,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Park
//		Park a channel.
//
func Park(socket *Socket, actionID, channel1, channel2 string, timeout int, parkinglot string) (map[string]string, error) {
	// verify action ID
	if len(channel1) == 0 || len(channel2) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Park",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel1,
		"\r\nChannel2: ",
		channel2,
		"\r\nTimeout: ",
		strconv.Itoa(timeout),
		"\r\nParkinglot: ",
		parkinglot,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	ParkedCalls
//		List parked calls.
//
func ParkedCalls(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("ParkedCalls", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"ParkedCall",
		"ParkedCallsComplete")
}

//
//	PlayDTMF
//		Play DTMF signal on a specific channel.
//
func PlayDTMF(socket *Socket, actionID, channel, digit string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(digit) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: PlayDTMF",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nDigit: ",
		digit,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Redirect
//		Redirect (transfer) a call.
//
func Redirect(socket *Socket, actionID, channel, exten, context, priority string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(exten) == 0 ||
		len(context) == 0 || len(priority) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: Redirect",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nExten: ",
		exten,
		"\r\nContext: ",
		context,
		"\r\nPriority: ",
		priority,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	SendText
//		Send text message to channel.
//
func SendText(socket *Socket, actionID, channel, msg string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(msg) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SendText",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nMessage: ",
		msg,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)

}

//
//	Setvar
//		Set a channel variable.
//		Set a global or local channel variable.
//		Note:	If a channel name is not provided then the variable is global.
//
func Setvar(socket *Socket, actionID, channel, variable, value string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(variable) == 0 || len(value) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Setvar",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nVariable: ",
		variable,
		"\r\nValue: ",
		value,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)

}

//
//	Status
//		List channel status.
//		Will return the status information of each channel along with the value for the specified channel variables.
//
func Status(socket *Socket, actionID, channel, variables string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(variables) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Status",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nVariables: ",
		variables,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	AGI
//		Add an AGI command to execute by Async AGI.
//
func AGI(socket *Socket, actionID, channel, agiCommand, agiCommandID string) (map[string]string, error) {
	// verify channel and action ID
	if len(channel) == 0 || len(actionID) == 0 ||
		len(agiCommand) == 0 || len(agiCommandID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: AGI",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nCommand: ",
		agiCommand,
		"\r\nCommandID: ",
		agiCommandID,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	AOCMessage
//		Generate an Advice of Charge message on a channel.
//
func AOCMessage(socket *Socket, actionID string, aocData AOCData) (map[string]string, error) {
	// verify channel and action ID
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: AOCMessage",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		aocData.Channel,
		"\r\nChannelPrefix: ",
		aocData.ChannelPrefix,
		"\r\nMsgType: ",
		aocData.MsgType,
		"\r\nChargeType: ",
		aocData.ChargeType,
		"\r\nUnitAmount(0): ",
		aocData.UnitAmount,
		"\r\nUnitType(0): ",
		aocData.UnitType,
		"\r\nCurrencyName: ",
		aocData.CurrencyName,
		"\r\nCurrencyAmount: ",
		aocData.CurrencyAmount,
		"\r\nCurrencyMultiplier: ",
		aocData.CurrencyMultiplier,
		"\r\nTotalType: ",
		aocData.TotalType,
		"\r\nAOCBillingId: ",
		aocData.AocBillingId,
		"\r\nChargingAssociationId: ",
		aocData.ChargingAssociationId,
		"\r\nChargingAssociationNumber: ",
		aocData.ChargingAssociationNumber,
		"\r\nChargingAssociationPlan: ",
		aocData.ChargingAssociationPlan,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Getvar
//		Gets a channel variable.
//
func Getvar(socket *Socket, actionID, channel, variable string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 || len(variable) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: Getvar",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nVariable: ",
		variable,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	LocalOptimizeAway
//		Optimize away a local channel when possible.
//		A local channel created with "/n" will not automatically optimize away.
//		Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
func LocalOptimizeAway(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: LocalOptimizeAway",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
