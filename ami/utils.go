package ami

import (
	"errors"
	"strings"
)

var (
	// ErrInvalidMessageParameters occurs when the parameters of message are invalid.
	ErrInvalidMessageParameters = errors.New("Invalid message parameters")
	// ErrInvalidActionID occurs when the action ID is invalid.
	ErrInvalidActionID = errors.New("Invalid Action ID")
	// ErrInvalidAction occurs when the action type is invalid.
	ErrInvalidAction = errors.New("invalid Action")
)

func decode(socket *Socket) (map[string]string, error) {
	message := make(map[string]string)
	for {
		s, err := socket.Recv()
		if err != nil {
			return nil, err
		}
		lines := strings.Split(s, "\r\n")
		for _, line := range lines {
			keys := strings.SplitAfterN(line, ":", 2)
			if len(keys) == 2 {
				action := strings.TrimSpace(strings.Trim(keys[0], ":"))
				response := strings.TrimSpace(keys[1])
				message[action] = response
			} else if strings.Contains(line, "\r\n\r\n") {
				return message, nil
			}
		}
	}
}

func sendCommand(socket *Socket, parameters map[string]string) (map[string]string, error) {
	var cmd string
	for key, value := range parameters {
		cmd += strings.TrimSpace(key) + ": \r\n" + strings.TrimSpace(value) + "\r\n"
	}
	cmd += "\r\n"
	if err := socket.Send(cmd); err != nil {
		return nil, err
	}
	resp, err := decode(socket)
	if err != nil || resp["ActionID"] != parameters["ActionID"] {
		return nil, err
	}
	return resp, nil
}
