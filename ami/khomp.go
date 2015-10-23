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
	errInvalidSMSParameters  = errors.New("khomp: Invalid SMS parameters")
	errInvalidSMSDevice      = errors.New("khomp: Invalid SMS device")
	errInvalidSMSDestination = errors.New("khomp: Invalid SMS destination")
	errInvalidSMSMessage     = errors.New("khomp: Invalid SMS message")
)

func checkKhompSMSData(data KhompSMSData) error {
	if len(data.Device) == 0 {
		return errInvalidSMSDevice
	}
	if len(data.Destination) == 0 {
		return errInvalidSMSDestination
	}
	if len(data.Message) == 0 {
		return errInvalidSMSMessage
	}
	return nil
}

//	KSendSMS
//		Send a SMS using KHOMP device
//
func KSendSMS(socket *Socket, actionID string, data KhompSMSData) (map[string]string, error) {

	if len(actionID) == 0 {
		return nil, errInvalidSMSParameters
	}
	if err := checkKhompSMSData(data); err != nil {
		return nil, err
	}

	s := map[bool]string{false: "false", true: "true"}
	command := []string{
		"Action: KSendSMS",
		"\r\nActionID: ",
		actionID,
		"\r\nDevice: ",
		data.Device,
		"\r\nDestination: ",
		data.Destination,
		"\r\nConfirmation: ",
		s[data.Confirmation],
		"\r\nMessage: ",
		data.Message,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
