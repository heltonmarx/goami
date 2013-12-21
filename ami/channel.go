package ami

import (
	"errors"
	"strconv"
)

//	AbsoluteTimeout	
//		Set absolute timeout.
//		Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
//
func SIPnotify(socket *Socket, actionID string, channel string, timeout int) (map[string]string, error) {
	//verify socket
	if !socket.Connected() {
		return nil, errors.New("Invalid socket")
	}

	// verify channel and action ID
	if len(channel) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: AbsoluteTimeout",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nTimeout: ",
		strconv.Itoa(timeout),
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
