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
	if action == "" {
		return nil, ErrInvalidAction
	}
	b, err := marshal(&struct {
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

func requestList(client Client, action, id, event, complete string) ([]Response, error) {
	if action == "" {
		return nil, ErrInvalidAction
	}
	b, err := marshal(&struct {
		Action string `ami:"Action"`
		ID     string `ami:"ActionID, omitempty"`
	}{Action: action, ID: id})
	if err != nil {
		return nil, err
	}
	if err := client.Send(string(b)); err != nil {
		return nil, err
	}
	for {
		input, err := client.Recv()
		if err != nil {
			return nil, err
		}
		responses, finish, err := parseEvent(event, complete, []string{input})
		if err != nil {
			return nil, err
		}
		if finish {
			return responses, nil
		}
	}
	return nil, fmt.Errorf("could not parse request %s", action)
}

func parseEvent(event, complete string, input []string) ([]Response, bool, error) {
	var list []Response
	verify := false
	for _, in := range input {
		rsp, err := parseResponse(in)
		if err != nil {
			return nil, false, err
		}
		if !verify {
			if success := rsp.Get("Response"); success != "Success" {
				return nil, false, fmt.Errorf("failed on event %s:%v\n", event, rsp.Get("Message"))
			}
			verify = true
		} else {
			evt := rsp.Get("Event")
			if evt == complete {
				return list, true, nil
			} else if evt == event {
				list = append(list, rsp)
			}
		}
	}
	return list, false, nil
}
