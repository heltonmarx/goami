package main

import (
	"fmt"
	"log"

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
	//Login
	uuid, _ := ami.GetUUID()
	if err := ami.Login(socket, "admin", "admin", "Off", uuid); err != nil {
		log.Fatalf("login error (%v)\n", err)
	}
	defer ami.Logoff(socket, uuid)
	fmt.Printf("login ok!\n")

	//List All SIPPeers
	peers, err := ami.SIPPeers(socket, uuid)
	if err != nil {
		log.Fatalf("sip peers error: %v\n", err)
	}
	for _, peer := range peers {
		message, err := ami.SIPShowPeer(socket, uuid, peer.Get("ObjectName"))
		if err != nil {
			log.Fatalf("sip show peer error: %v\n", err)
		}
		log.Printf("message: [%q]\n", message)
	}
}
