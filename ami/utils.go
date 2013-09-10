package ami

import (
	"strings"
)

type Answer struct {
	action   string
	response string
}

func parseAnswer(socket *Socket) ([]Answer, error) {
	answers := make([]Answer, 0)

	for {
		s, err := socket.Recv()
		if err != nil {
			return answers, err
		}
		line := strings.Split(s, "\r\n")
		for i := 0; i < len(line); i++ {
			keys := strings.Split(line[i], ":")
			if len(keys) == 2 {
				action := strings.TrimSpace(keys[0])
				response := strings.TrimSpace(keys[1])
				answers = append(answers, Answer{
					action:   action,
					response: response,
				})
			} else {
				return answers, nil
			}
		}
	}
	return answers, nil
}

func getResponse(answers[]Answer, action string) (string) {
	for i:= 0; i < len(answers); i++ {
		if answers[i].action == action {
			return answers[i].response
		}
	}
	return ""
}

func cmpActionID(answers[]Answer, actionID string) (bool) {
	for i:= 0; i < len(answers); i++ {
		if (answers[i].action == "ActionID") && (answers[i].response == actionID) {
			return true
		}
	}
	return false

}
