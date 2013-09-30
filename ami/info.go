package ami

import (
	"errors"
)

type opCode int

const (
	peerGetResponse opCode = iota
	peerGetList
)

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
	state := peerGetResponse
	for {
		message, err := decode(socket)
		if (err != nil) || (message["ActionID"] != actionID) {
			return nil, err
		}
		switch state {
		case peerGetResponse:
			if message["Response"] != "Success" {
				return nil, errors.New(message["Message"])
			} else {
				state = peerGetList
			}
		case peerGetList:
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
