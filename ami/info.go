package ami

import (
	"errors"
)

type SIPPeer struct {
	Channeltype    string
	ObjectName     string
	ChanObjectType string
	IPaddress      string
	IPport         int
	Dynamic        bool
	Forceport      bool
	VideoSupport   bool
	TextSupport    bool
	ACL            bool
	RealtimeDevice bool
	Status         string
}

type opCode int

const (
	peerGetResponse opCode = iota
	peerGetList
)

func SIPPeers(socket *Socket, actionID string) ([]SIPPeer, error) {
	var err error
	var state opCode
	var p SIPPeer

	if !socket.Connected() {
		return nil, errors.New("Invalid socket")
	}

	peerCmd := []string{
		"Action: SIPpeers",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err = sendCmd(socket, peerCmd)
	if err != nil {
		return nil, err
	}

	/* set state to initial state */
	list := make([]SIPPeer, 0)
	state = peerGetResponse
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
				unmarshal(&p, message)
				list = append(list, p)
			}
		}
	}
on_exit:
	return list, nil
}
