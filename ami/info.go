package ami

import (
	"errors"
	"fmt"
)

type opCode int

const (
	getResponse opCode = iota
	getList
)

func getList(socket *Socket, action, actionID, event, complete string ) ([]map[string]string, error) {

	if !socket.Connected() {
		return nil, errors.New("Invalid socket")
	}

	// verify action ID
	if len(actionID) == 0 || len(action) == 0 ||
        len(event) == 0 || len(complete) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: ",
        action,
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}

	err := sendCmd(socket, command)
	if err != nil {
		return nil, err
	}

	list := make([]map[string]string, 0)
	state := getResponse
	for {
		message, err := decode(socket)
		if (err != nil) || (message["ActionID"] != actionID) {
			return nil, err
		}
		switch state {
		case getResponse:
			if message["Response"] != "Success" {
				return nil, errors.New(message["Message"])
			} else {
				state = getList
			}
		case getList:
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


func SIPPeers(socket *Socket, actionID string) ([]map[string]string, error) {
	var err error

	if !socket.Connected() {
		return nil, errors.New("Invalid socket")
	}

	// verify parameters
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SIPpeers",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err = sendCmd(socket, command)
	if err != nil {
		return nil, err
	}

	/* set state to initial state */
	list := make([]map[string]string, 0)
	state := getResponse
	for {
		message, err := decode(socket)
		if (err != nil) || (message["ActionID"] != actionID) {
			return nil, err
		}
		switch state {
		case getResponse:
			if message["Response"] != "Success" {
				return nil, errors.New(message["Message"])
			} else {
				state = getList
			}
		case getList:
			if message["Event"] == "PeerlistComplete" {
				goto on_exit
			} else if message["Event"] == "PeerEntry" {
				list = append(list, message)
			}
		}
	}
on_exit:
	return list, nil
}

func SIPShowpeer(socket *Socket, actionID string, name string) (map[string]string, error) {

	if !socket.Connected() {
		return nil, errors.New("Invalid socket")
	}

	// verify name and action ID
	if len(name) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SIPshowpeer",
		"\r\nActionID: ",
		actionID,
		"\r\nPeer: ",
		name,
		"\r\n\r\n", // end of command
	}

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

//  Agents
//  Lists agents and their status.
//
func Agents(socket *Socket, actionID string) ([]map[string]string, error) {
    return getList(socket, "Agents", actionID, "AgentsEntry", "AgentsComplete")
}

