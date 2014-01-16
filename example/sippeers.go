// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package main

import (
	"fmt"
	"github.com/heltonmarx/goami/ami"
)

func main() {
	socket, err := ami.NewSocket("127.0.0.1:5038")
	if err != nil {
		fmt.Printf("socket error: %v\n", err)
		return
	}
	if _, err := ami.Connect(socket); err != nil {
		return
	}
	var ret bool

	//Login
	uuid, _ := ami.GetUUID()
	ret, err = ami.Login(socket, "admin", "admin", "Off", uuid)
	if err != nil || ret == false {
		fmt.Printf("login error (%v)\n", err)
		return
	}
	fmt.Printf("login ok!\n")

	//List All SIPPeers
	list, _ := ami.SIPpeers(socket, uuid)
	if len(list) > 0 {
		for _, m := range list {
			message, _ := ami.SIPshowpeer(socket, uuid, m["ObjectName"])
			for k, v := range message {
				fmt.Printf("%s : %q\n", k, v)
			}
		}
	}
	fmt.Printf("\n")

	//Logoff
	fmt.Printf("logoff\n")
	ret, err = ami.Logoff(socket, uuid)
	if err != nil || ret == false {
		fmt.Printf("logoff error: (%v)\n", err)
		return
	}
	fmt.Printf("goodbye !\n")
}
