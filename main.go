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
	list, _ := ami.SIPPeers(socket, uuid)
	if len(list) > 0 {
		for _, m := range list {
			message, _ := ami.SIPShowpeer(socket, uuid, m["ObjectName"])
			for k, v := range message {
				fmt.Printf("%s : %q\n", k, v)
			}
		}
	}
	fmt.Printf("\n")

	/**
	 *	Agents
	 */
	agents, _ := ami.Agents(socket, uuid)
	fmt.Printf("\n\nagents: %d\n", len(agents))
	for i, j := range agents {
		fmt.Printf("%s: %q\n", i, j)
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
