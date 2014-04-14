// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
)

//  SKINNYdevices
//		List SKINNY devices (text format).
//		Lists Skinny devices in text format with details on current status.
//		Devicelist will follow as separate events,
//		followed by a final event called DevicelistComplete.
//
func SKINNYdevices(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("SKINNYdevices", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"DeviceEntry", "DevicelistComplete")
}

//	SKINNYlines
//		List SKINNY lines (text format).
//		Lists Skinny lines in text format with details on current status.
//		Linelist will follow as separate events,
//		followed by a final event called LinelistComplete.
//
func SKINNYlines(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("SKINNYdevices", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"LineEntry", "LinelistComplete")
}

//	SKINNYshowdevice
//		Show SKINNY device (text format).
//		Show one SKINNY device with details on current status.
//
func SKINNYshowdevice(socket *Socket, actionID, device string) (map[string]string, error) {
	if len(actionID) == 0 || len(device) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: SKINNYshowdevice",
		"\r\nActionID: ",
		actionID,
		"\r\nDevice: ",
		device,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	SKINNYshowline
//		Show SKINNY line (text format).
//		Show one SKINNY line with details on current status.
//
func SKINNYshowline(socket *Socket, actionID, line string) (map[string]string, error) {
	if len(actionID) == 0 || len(line) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: SKINNYshowline",
		"\r\nActionID: ",
		actionID,
		"\r\nLine: ",
		line,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
