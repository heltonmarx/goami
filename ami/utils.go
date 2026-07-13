package ami

import (
	"bytes"
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
)

// GetUUID returns a new UUID string.
func GetUUID() (string, error) {
	return uuid.New().String(), nil
}

func command(action string, id string, v ...any) ([]byte, error) {
	if action == "" {
		return nil, errors.New("invalid Action")
	}
	return marshal(&struct {
		Action string `ami:"Action"`
		ID     string `ami:"ActionID, omitempty"`
		V      []any
	}{Action: action, ID: id, V: v})
}

func send(ctx context.Context, client Client, action, id string, v any) (Response, error) {
	b, err := command(action, id, v)
	if err != nil {
		return nil, err
	}
	if err := client.Send(string(b)); err != nil {
		return nil, err
	}
	return read(ctx, client)
}

func read(ctx context.Context, client Client) (Response, error) {
	var buffer bytes.Buffer
	for {
		input, err := client.Recv(ctx)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(input)
		if strings.HasSuffix(buffer.String(), "\r\n\r\n") {
			break
		}
	}
	return parseResponse(buffer.String())
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
		} else if strings.Contains(line, "\r\n\r\n") || line == "" {
			return resp, nil
		}
	}
	return resp, nil
}

func requestList(ctx context.Context, client Client, action, id, event, complete string, v ...any) ([]Response, error) {
	b, err := command(action, id, v)
	if err != nil {
		return nil, err
	}
	if err := client.Send(string(b)); err != nil {
		return nil, err
	}

	response := make([]Response, 0)
	for {
		rsp, err := read(ctx, client)
		if err != nil {
			return nil, err
		}
		e := rsp.Get("Event")
		r := rsp.Get("Response")
		if e == event {
			response = append(response, rsp)
		} else if e == complete || r != "" && r != "Success" {
			break
		}
	}
	return response, nil
}

// requestMultiEvent allows for a list of events to be specified, used in cases where a command
// returns multiple types of events.
func requestMultiEvent(ctx context.Context, client Client, action, id string, events []string, complete string, v ...any) ([]Response, error) {
	set := make(map[string]struct{}, len(events))
	for _, evt := range events {
		set[evt] = struct{}{}
	}

	b, err := command(action, id, v)
	if err != nil {
		return nil, err
	}
	if err := client.Send(string(b)); err != nil {
		return nil, err
	}

	response := make([]Response, 0)
	for {
		rsp, err := read(ctx, client)
		if err != nil {
			return nil, err
		}
		e := rsp.Get("Event")
		r := rsp.Get("Response")

		_, ok := set[e]
		if ok {
			response = append(response, rsp)
		} else if e == complete || r != "" && r != "Success" {
			break
		}
	}
	return response, nil
}
