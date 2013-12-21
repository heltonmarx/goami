package ami

import (
	"errors"
)

const (
	getResponseState int = iota
	getListState
)

func getMessageList(socket *Socket, action, actionID, event, complete string) ([]map[string]string, error) {
	// verify socket
	if !socket.Connected() {
		return nil, errors.New("Invalid socket")
	}

	// verify parameters
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
	state := getResponseState
	for {
		message, err := decode(socket)
		if (err != nil) || (message["ActionID"] != actionID) {
			return nil, err
		}
		switch state {
		case getResponseState:
			if message["Response"] != "Success" {
				return nil, errors.New(message["Message"])
			} else {
				state = getListState
			}
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

//  SIPPeers
//      Lists SIP peers in text format with details on current status. 
//      Peerlist will follow as separate events, followed by a final event called PeerlistComplete
//
func SIPPeers(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket, "SIPpeers", actionID, "PeerEntry", "PeerlistComplete")
}

//  SIPShowpeer
//      Show one SIP peer with details on current status.
//
func SIPShowpeer(socket *Socket, actionID string, name string) (map[string]string, error) {

	//verify socket
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
//      Lists agents and their status.
//
func Agents(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket, "Agents", actionID, "AgentsEntry", "AgentsComplete")
}
