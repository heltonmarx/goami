package main

import (
	"./ami"
	"fmt"
)

func iax(socket *ami.Socket, actionID string) {
	iaxPeerList, _ := ami.IAXpeerlist(socket, actionID)
	if len(iaxPeerList) > 0 {
		for _, m := range iaxPeerList {
			fmt.Printf("IAX: %v\n", m)
		}
	}

	iaxPeer, _ := ami.IAXpeers(socket, actionID)
	if len(iaxPeer) > 0 {
		for _, m := range iaxPeer {
			fmt.Printf("IAX2: %v\n", m)
		}
	}
	iaxRegistry, _ := ami.IAXregistry(socket, actionID)
	if len(iaxRegistry) > 0 {
		for _, m := range iaxPeer {
			fmt.Printf("IAX Registry: %v\n", m)
		}
	}
}

func mailbox(socket *ami.Socket, actionID string) {
	mc, _ := ami.MailboxCount(socket, actionID, "mailbox@vm-context")
	if mc != nil {
		fmt.Printf("MailboxCount: %v\n", mc)
	}

	ms, _ := ami.MailboxStatus(socket, actionID, "mailbox@vm-context")
	if ms != nil {
		fmt.Printf("Mailbox status: %v\n", ms)
	}
}

func main() {
	fmt.Printf("goami %s\n", ami.VersionInfo())

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

	/**
	 *	Agents
	 */
	agents, _ := ami.Agents(socket, uuid)
	fmt.Printf("agents: %d\n", len(agents))
	for i, j := range agents {
		fmt.Printf("%s: %q\n", i, j)
	}
	fmt.Printf("\n")

	/**
	 *	IAX Functions
	 */
	iax(socket, uuid)

	/**
	 *	Mailbox Functions
	 */
	mailbox(socket, uuid)

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
