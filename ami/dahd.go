// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
)

var (
	errInvalidDAHDIDParameters = errors.New("DHADID: Invalid parameters")
)

//	DAHDIDialOffhook
//		Dial over DAHDI channel while offhook.
//		Generate DTMF control frames to the bridged peer.
//
func DAHDIDialOffhook(socket *Socket, actionID, channel, number string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 || len(number) == 0 {
		return nil, errInvalidDAHDIDParameters
	}
	command := []string{
		"Action: DAHDIDialOffhook",
		"\r\nActionID: ",
		actionID,
		"\r\nDAHDIChannel: ",
		channel,
		"\r\nNumber: ",
		number,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	DAHDIDNDoff
//		Toggle DAHDI channel Do Not Disturb status OFF.
//
func DAHDIDNDoff(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errInvalidDAHDIDParameters
	}
	command := []string{
		"Action: DAHDIDNDoff",
		"\r\nActionID: ",
		actionID,
		"\r\nDAHDIChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

// DAHDIDNDon
//		Toggle DAHDI channel Do Not Disturb status ON.
//
func DAHDIDNDon(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errInvalidDAHDIDParameters
	}
	command := []string{
		"Action: DAHDIDNDon",
		"\r\nActionID: ",
		actionID,
		"\r\nDAHDIChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	DAHDIHangup
//		Hangup DAHDI Channel.
//
func DAHDIHangup(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errInvalidDAHDIDParameters
	}
	command := []string{
		"Action: DAHDIHangup",
		"\r\nActionID: ",
		actionID,
		"\r\nDAHDIChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	DAHDIRestart
//		Fully Restart DAHDI channels (terminates calls).
//
func DAHDIRestart(socket *Socket, actionID string) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errInvalidDAHDIDParameters
	}
	command := []string{
		"Action: DAHDIRestart",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	DAHDIShowChannels
//		Show status of DAHDI channels.
//
func DAHDIShowChannels(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errInvalidDAHDIDParameters
	}
	command := []string{
		"Action: DAHDIShowChannels",
		"\r\nActionID: ",
		actionID,
		"\r\nDAHDIChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	DAHDITransfer
//		Transfer DAHDI Channel.
//
func DAHDITransfer(socket *Socket, actionID, channel string) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 {
		return nil, errInvalidDAHDIDParameters
	}
	command := []string{
		"Action: DAHDITransfer",
		"\r\nActionID: ",
		actionID,
		"\r\nDAHDIChannel: ",
		channel,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
