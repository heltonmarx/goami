// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
)

type Socket struct {
	conn    net.Conn
	buffer  *bufio.Reader
	address string
}

var (
	ErrNotConnected = errors.New("socket: net.Conn not connected")
)

func (socket *Socket) Connected() bool {
	if socket.conn == nil {
		return false
	}
	return true
}

func (socket *Socket) Disconnect() error {
	if socket.conn != nil {
		return socket.conn.Close()
	}
	return nil
}

func (socket *Socket) Connect() error {
	if socket.Connected() {
		socket.Disconnect()
	}
	var err error
	socket.conn, err = net.Dial("tcp", socket.address)

	// connected succesfull
	if err == nil {
		socket.buffer = bufio.NewReaderSize(socket.conn, 8192)
	}
	return err
}

func (socket *Socket) Send(format string, a ...interface{}) error {
	if !socket.Connected() {
		return ErrNotConnected
	}
	m := fmt.Sprintf(format, a...)
	fmt.Fprint(socket.conn, m)
	return nil
}

func (socket *Socket) Recv() (string, error) {
	if !socket.Connected() {
		return "", ErrNotConnected
	}
	buf := make([]byte, 0)
	for {
		b, err := socket.buffer.ReadBytes('\n')
		if err != nil {
			return "", err
		}
		buf = append(buf, b...)
		if len(bytes.TrimSpace(b)) == 0 &&
			strings.Contains(string(buf), string('\n')) == true {
			break
		} else if socket.buffer.Buffered() == 0 {
			return string(buf), nil
		}
	}
	return string(buf), nil
}

func NewSocket(address string) (*Socket, error) {
	socket := Socket{address: address}
	if err := socket.Connect(); err != nil {
		return nil, err
	}
	return &socket, nil
}
