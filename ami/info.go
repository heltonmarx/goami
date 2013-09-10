package ami

import (
	"errors"
	"fmt"
	"strconv"
)

type SIPPeer struct {
	Channeltype    string
	ObjectName     string
	ChanObjectType string
	IPaddress      string
	IPport         int
	Dynamic        int
	Forceport      int
	VideoSupport   int
	TextSupport    int
	ACL            int
	RealtimeDevice int
	Status         string
}

type opCode int

const (
	peerGetResponse opCode = iota
	peerGetList
)

func SIPPeers(socket *Socket, actionID string) (int, error) {
	if !socket.Connected() {
		return 0, errors.New("Invalid socket")
	}
	var err error
	var response string
	var state opCode

	peerCmd := []string{
		"Action: SIPpeers",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err = sendCmd(socket, peerCmd)
	if err != nil {
		return 0, err
	}

	/* set state to initial state */
	state = peerGetResponse
	for {
		answers, err := parseAnswer(socket)
		if (err != nil) || (cmpActionID(answers, actionID) == false) {
			return 0, err
		}
		fmt.Printf("answers: %q\n", answers)
		switch state {
		case peerGetResponse:
			response = getResponse(answers, "Response")
			if response != "Success" {
				response = getResponse(answers, "Message")
				return 0, errors.New(response)
			} else {
				state = peerGetList
			}
		case peerGetList :
			response = getResponse(answers, "Event")
			if response == "PeerlistComplete" {
				n := getResponse(answers,"ListItems")
				return strconv.Atoi(n) 

			}
		}
	}
	return 0, nil
}
