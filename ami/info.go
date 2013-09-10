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

func parseSIPPeers(answers []Answer) SIPPeer {
	var p SIPPeer
	for i := 0; i < len(answers); i++ {
		switch answers[i].action {
		case "Channeltype":
			p.Channeltype = answers[i].response
		case "ObjectName":
			p.ObjectName = answers[i].response
		case "ChanObjectType":
			p.ChanObjectType = answers[i].response
		case "IPaddress":
			p.IPaddress = answers[i].response
		case "Status":
			p.Status = answers[i].response
		// integer value
		case "IPport":
			p.IPport, _ = strconv.Atoi(answers[i].response)
		//boolean values
		case "Dynamic":
			p.Dynamic = parseBool(answers[i].response)
		case "Forceport":
			p.Forceport = parseBool(answers[i].response)
		case "VideoSupport":
			p.VideoSupport = parseBool(answers[i].response)
		case "TextSupport":
			p.TextSupport = parseBool(answers[i].response)
		case "ACL":
			p.ACL = parseBool(answers[i].response)
		case "RealtimeDevice":
			p.RealtimeDevice = parseBool(answers[i].response)
		}
	}
	return p
}

func SIPPeers(socket *Socket, actionID string) ([]SIPPeer, error) {
	var err error
	var response string
	var state opCode

	sippeer := make([]SIPPeer, 0)

	if !socket.Connected() {
		return sippeer, errors.New("Invalid socket")
	}

	peerCmd := []string{
		"Action: SIPpeers",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err = sendCmd(socket, peerCmd)
	if err != nil {
		return sippeer, err
	}

	/* set state to initial state */
	state = peerGetResponse
	for {
		answers, err := parseAnswer(socket)
		if (err != nil) || (cmpActionID(answers, actionID) == false) {
			return sippeer, err
		}
		switch state {
		case peerGetResponse:
			response = getResponse(answers, "Response")
			if response != "Success" {
				response = getResponse(answers, "Message")
				return sippeer, errors.New(response)
			} else {
				state = peerGetList
			}
		case peerGetList:
			response = getResponse(answers, "Event")
			if response == "PeerlistComplete" {
				return sippeer, nil

			} else if response == "PeerEntry" {
				//decoding and append SIPPeer
				sippeer = append(sippeer, parseSIPPeers(answers))
			}
		}
	}
	return sippeer, nil
}
