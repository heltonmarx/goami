package ami

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrInvalidAction occurs when the action type is invalid.
	ErrInvalidAction = errors.New("invalid Action")
)

func send(client Client, action, id string, v interface{}) (Response, error) {
	var err error
	var b []byte

	if action == "" {
		return nil, ErrInvalidAction
	}
	b, err = marshal(&struct {
		Action string `ami:"Action"`
		ID     string `ami:"ActionID, omitempty"`
		V      interface{}
	}{Action: action, ID: id, V: v})
	if err != nil {
		return nil, err
	}
	if err := client.Send(string(b)); err != nil {
		return nil, err
	}
	input, err := client.Recv()
	if err != nil {
		return nil, err
	}
	return parseResponse(input)
}

/*
func requestEvent(client Client, action, id, event, complete string) ([]Response, error) {
	cmd := newCommand(action, id)
	if err := client.Send(cmd.String()); err != nil {
		return nil, err
	}
}
*/

func parseResponse(input string) (Response, error) {
	resp := make(Response)

	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		keys := strings.SplitAfterN(line, ":", 2)
		if len(keys) == 2 {
			key := strings.TrimSpace(strings.Trim(keys[0], ":"))
			value := strings.TrimSpace(keys[1])
			resp[key] = append(resp[key], value)
		} else if strings.Contains(line, "\r\n\r\n") || len(line) == 0 {
			return resp, nil
		}
	}
	return resp, nil
}

func parseEvent(event, complete string, input []string) ([]Response, error) {
	list := make([]Response, 0)
	verify := false
	for _, in := range input {
		rsp, err := parseResponse(in)
		if err != nil {
			return nil, err
		}
		if !verify {
			if success := rsp.Get("Response"); success != "Success" {
				return nil, fmt.Errorf("failed on event %s:%v\n", event, rsp.Get("Message"))
			}
			verify = true
		} else {
			evt := rsp.Get("Event")
			if evt == complete {
				return list, nil
			} else if evt == event {
				list = append(list, rsp)
			}
		}
	}
	return list, nil
}
