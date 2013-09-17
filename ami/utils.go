package ami

import (
	"strings"
)

func parseMessage(socket *Socket) (map[string]string, error) {
	message := make(map[string]string, 0)

	for {
		s, err := socket.Recv()
		if err != nil {
			return nil, err
		}
		line := strings.Split(s, "\r\n")
		for i := 0; i < len(line); i++ {
			keys := strings.Split(line[i], ":")
			if len(keys) == 2 {
				action := strings.TrimSpace(keys[0])
				response := strings.TrimSpace(keys[1])
				message[action] = response
			} else if strings.Contains(s, "\r\n\r\n") {
				goto on_exit
			}
		}
	}
on_exit:
	return message, nil
}
