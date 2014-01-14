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
)

type Socket struct {
	conn    net.Conn
	buffer  *bufio.Reader
	address string
}

func (self *Socket) Connected() bool {
	if self.conn == nil {
		return false
	}
	return true
}

func (self *Socket) Disconnect() error {
	if self.conn != nil {
		return self.conn.Close()
	}
	return nil
}

func (self *Socket) Connect() error {
	if self.Connected() {
		self.Disconnect()
	}
	var err error
	self.conn, err = net.Dial("tcp", self.address)

	// connected succesfull
	if err == nil {
		self.buffer = bufio.NewReaderSize(self.conn, 8192)
	}
	return err
}

func (self *Socket) Send(format string, a ...interface{}) error {
	if !self.Connected() {
		return errors.New("Not connected to AMI\n")
	}
	m := fmt.Sprintf(format, a...)
	fmt.Fprint(self.conn, m)
	return nil
}

func (self *Socket) Recv() (string, error) {
	bytesRead := make([]byte, 0)
	var readLine []byte
	var err error

	for {
		readLine, err = self.buffer.ReadBytes('\n')
		if err != nil {
			return "", nil
		}
		bytesRead = append(bytesRead, readLine...)
		if len(bytes.TrimSpace(readLine)) == 0 {
			break
		} else if self.buffer.Buffered() == 0 {
			return string(bytesRead), nil
		}
	}
	return string(bytesRead), nil
}

func NewSocket(address string) (*Socket, error) {
	socket := Socket{
		address: address,
	}
	err := socket.Connect()
	if err != nil {
		return nil, err
	}
	return &socket, nil
}
