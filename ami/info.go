package ami

import (
	"errors"
	"fmt"
	"strings"
)

/**
  first type of message:
  Test #1:
      Response: Success
      EventList: start
      Message: Peer status list will follow

      Event: PeerlistComplete
      EventList: Complete
      ListItems: 0

  Test #2:
      Response: Success
      EventList: start
      Message: Peer status list will follow

      Event: PeerEntry
      Channeltype: SIP
      ObjectName: 1000
      ChanObjectType: peer
      IPaddress: -none-
      IPport: 0
      Dynamic: yes
      Forcerport: yes
      VideoSupport: no
      TextSupport: no
      ACL: yes
      Status: Unmonitored
      RealtimeDevice: no

      Event: PeerEntry
      Channeltype: SIP
      ObjectName: 1001
      ChanObjectType: peer
      IPaddress: -none-
      IPport: 0
      Dynamic: yes
      Forcerport: yes
      VideoSupport: no
      TextSupport: no
      ACL: yes
      Status: Unmonitored
      RealtimeDevice: no

      Event: PeerEntry
      Channeltype: SIP
      ObjectName: fooprovider
      ChanObjectType: peer
      IPaddress: -none-
      IPport: 0
      Dynamic: no
      Forcerport: yes
      VideoSupport: no
      TextSupport: no
      ACL: no
      Status: Unmonitored
      RealtimeDevice: no

      Event: PeerlistComplete
      EventList: Complete
      ListItems: 3

*/
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

func SIPPeers(socket *Socket, actionID string) (string, error) {
	if !socket.Connected() {
		return "", errors.New("Invalid socket")
	}
	var answer string
	var err error

	peerCmd := []string{
		"Action: SIPpeers",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err = sendCmd(socket, peerCmd)
	if err != nil {
		return "", err
	}
	answer, err = socket.Recv()
	fmt.Printf("answer: %v\n", answer)
	if err != nil || !strings.Contains(answer, "Success") {
		return "", errors.New("SIPPeers failed")
	}
	answer, err = socket.Recv()
	if err != nil || !strings.Contains(answer, "PeerlistComplete") {
		return "", errors.New("SIPPeers failed")
	}
	return answer, err
}
