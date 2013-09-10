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

/*
	{"Event" "PeerEntry"} {"ActionID" "7f0a3ad2-6bd2-3a2f-b209-f3acb01b024e"} 
	{"Channeltype" "SIP"} {"ObjectName" "fooprovider"} {"ChanObjectType" "peer"} 
	{"IPaddress" "-none-"} {"IPport" "0"} {"Dynamic" "no"} {"Forcerport" "yes"} 
	{"VideoSupport" "no"} {"TextSupport" "no"} {"ACL" "no"} 
	{"Status" "Unmonitored"} {"RealtimeDevice" "no"}

*/
func parseBool(s string) bool {
	if s == "no" {
		return false
	}
	return true
}

func parseSIPPeers(answers []Answer) SIPPeer {
	var p SIPPeer

	p.Channeltype = getResponse(answers, "Channeltype")
	p.ObjectName = getResponse(answers, "ObjectName")
	p.ChanObjectType = getResponse(answers, "ChanObjectType")
	p.IPaddress = getResponse(answers, "IPaddress")
	p.IPport, _ = strconv.Atoi(getResponse(answers, "IPport"))
	p.Status = getResponse(answers, "Status")
	// boolean values 
	p.Dynamic = parseBool(getResponse(answers, "Dynamic"))
	p.Forceport = parseBool(getResponse(answers, "Forceport"))
	p.VideoSupport = parseBool(getResponse(answers, "VideoSupport"))
	p.TextSupport = parseBool(getResponse(answers, "TextSupport"))
	p.ACL = parseBool(getResponse(answers, "ACL"))
	p.RealtimeDevice = parseBool(getResponse(answers, "RealtimeDevice"))

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
