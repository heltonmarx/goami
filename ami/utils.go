package ami

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var mutex sync.Mutex

// GetUUID returns a new UUID based on /dev/urandom (unix).
func GetUUID() (string, error) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		return "", fmt.Errorf("open /dev/urandom error:[%v]", err)
	}
	defer f.Close()
	b := make([]byte, 16)

	_, err = f.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func command(action string, id string, v ...interface{}) ([]byte, error) {
	if action == "" {
		return nil, errors.New("invalid Action")
	}
	return marshal(&struct {
		Action string `ami:"Action"`
		ID     string `ami:"ActionID, omitempty"`
		V      []interface{}
	}{Action: action, ID: id, V: v})
}

func send(ctx context.Context, client Client, action, id string, v interface{}) (Response, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if id == "" {
		id, _ = GetUUID()
	}
	b, err := command(action, id, v)
	if err != nil {
		return nil, err
	}
	if err := client.Send(string(b)); err != nil {
		return nil, err
	}
	for {
		var response Response
		var err error
		response, err = read(ctx, client)
		if err != nil {
			return nil, err
		}
		actionID := response.Get("ActionID")
		if actionID == id {
			return response, nil
		}
	}
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

func requestList(ctx context.Context, client Client, action, id, event, complete string, v ...interface{}) ([]Response, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if id == "" {
		id, _ = GetUUID()
	}
	b, err := command(action, id, v)
	if err != nil {
		return nil, err
	}
	if err := client.Send(string(b)); err != nil {
		return nil, err
	}

	response := make([]Response, 0)
	// find the response associated to the command
	for {
		var rsp Response
		var err error
		rsp, err = read(ctx, client)
		if err != nil {
			return nil, err
		}
		actionID := rsp.Get("ActionID")
		if actionID != id {
			continue
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
