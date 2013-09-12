package ami

import (
	"errors"
	"strconv"
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

func parseBool(s string) bool {
	if s == "no" {
		return false
	}
	return true
}

func parseSIPPeers(message map[string]string) SIPPeer {
	var p SIPPeer

	p.Channeltype = message["Channeltype"]
	p.ObjectName = message["ObjectName"]
	p.ChanObjectType = message["ChanObjectType"]
	p.IPaddress = message["IPaddress"]
	p.Status = message["Status"]
	// integer value
	p.IPport, _ = strconv.Atoi(message["IPport"])
	//boolean values
	p.Dynamic = parseBool(message["Dynamic"])
	p.Forceport = parseBool(message["Forceport"])
	p.VideoSupport = parseBool(message["VideoSupport"])
	p.TextSupport = parseBool(message["TextSupport"])
	p.ACL = parseBool(message["ACL"])
	p.RealtimeDevice = parseBool(message["RealtimeDevice"])
	return p
}

func SIPPeers(socket *Socket, actionID string) ([]SIPPeer, error) {
	var err error
	var state opCode

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
	sippeer := make([]SIPPeer, 0)
	state = peerGetResponse
	for {
		message, err := parseMessage(socket)
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
				//decoding and append SIPPeer
				sippeer = append(sippeer, parseSIPPeers(message))
			}
		}
	}
on_exit:
	return sippeer, nil
}
