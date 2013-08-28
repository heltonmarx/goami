package ami

import (
    "errors"
    "strings"
    "fmt"
)

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
    answer, err = socket.Recv();
    fmt.Printf("answer: %v\n", answer);
	if err != nil || !strings.Contains(answer, "Success") {
        return "", errors.New("SIPPeers failed")
	}
    answer, err = socket.Recv();
	if err != nil || !strings.Contains(answer, "PeerlistComplete") {
        return "", errors.New("SIPPeers failed")
	}
    return answer, err
}
