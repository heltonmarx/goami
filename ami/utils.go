// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
	"strings"
)

var (
	ErrInvalidMessageParameters = errors.New("Invalid message parameters")
	ErrInvalidActionID          = errors.New("Invalid Action ID")
	ErrInvalidAction            = errors.New("invalid Action")
)

func decode(socket *Socket) (map[string]string, error) {
	message := make(map[string]string)

	for {
		s, err := socket.Recv()
		if err != nil {
			return nil, err
		}
		line := strings.Split(s, "\r\n")
		for i := 0; i < len(line); i++ {
			keys := strings.SplitAfterN(line[i], ":", 2)
			if len(keys) == 2 {
				action := strings.TrimSpace(strings.Trim(keys[0], ":"))
				response := strings.TrimSpace(keys[1])
				message[action] = response
			} else if strings.Contains(s, "\r\n\r\n") {
				goto on_exit
			} else {
				break
			}
		}
	}
on_exit:
	return message, nil
}

const (
	getResponseState int = iota
	getListState
)

func getMessageList(socket *Socket, command []string, actionID, event, complete string) ([]map[string]string, error) {
	// verify parameters
	if len(actionID) == 0 ||
		len(event) == 0 || len(complete) == 0 {
		return nil, ErrInvalidMessageParameters
	}

	err := sendCmd(socket, command)
	if err != nil {
		return nil, err
	}

	list := make([]map[string]string, 0)
	state := getResponseState
	for {
		message, err := decode(socket)
		//verify decode error
		if err != nil {
			return nil, err
		}

		//verify action id
		if len(message["ActionID"]) > 0 && (message["ActionID"] != actionID) {
			return nil, ErrInvalidActionID
		}

		switch state {
		case getResponseState:
			if message["Response"] != "Success" {
				return nil, errors.New(message["Message"])
			}
			state = getListState
		case getListState:
			if message["Event"] == complete {
				goto on_exit
			} else if message["Event"] == event {
				list = append(list, message)
			}
		}
	}
on_exit:
	return list, nil

}

func getMessage(socket *Socket, command []string, actionID string) (map[string]string, error) {
	err := sendCmd(socket, command)
	if err != nil {
		return nil, err
	}

	message, err := decode(socket)
	if (err != nil) || (message["ActionID"] != actionID) {
		return nil, err
	}
	return message, nil
}

func getCommand(action, actionID string) ([]string, error) {
	if len(action) == 0 {
		return nil, ErrInvalidAction
	}
	if len(actionID) == 0 {
		return nil, ErrInvalidActionID
	}
	command := []string{
		"Action: ",
		action,
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	return command, nil
}
