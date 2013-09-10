package main

import (
	"./ami"
	"fmt"
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

	/**
	 *  Login
	 */
	uuid, _ := ami.GetUUID()
	ret, err = ami.Login(socket, "admin", "admin", "Off", uuid)
	if err != nil || ret == false {
		fmt.Printf("login error (%v)\n", err)
		return
	}
	fmt.Printf("login ok!\n")

	/**
	 *  Ping
	 */
	ret, err = ami.Ping(socket, uuid)
	if err != nil || ret == false {
		fmt.Printf("ping error (%v)\n", err)
		return
	}

	/**
	 *  SIPPeers
	 */
	sippeer, _ := ami.SIPPeers(socket, uuid)
	fmt.Printf("sippeer qtd: %d\n", len(sippeer))
	for i := 0; i < len(sippeer); i++ {
		fmt.Printf("peer[%d]:[%s]\n", i, sippeer[i].ObjectName)
	}

	/**
	 *  Logoff
	 */
	fmt.Printf("logoff\n")
	ret, err = ami.Logoff(socket, uuid)
	if err != nil || ret == false {
		fmt.Printf("logoff error: (%v)\n", err)
		return
	}
	fmt.Printf("goodbye !\n")

}
