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
	list, err := ami.SIPPeers(socket, uuid)
	if err != nil {
		log.Fatalf("sip peers error: %v\n", err)
	}
	message, err := ami.SIPShowPeer(socket, uuid, list["ObjectName"])
	if err != nil {
		log.Fatalf("sip show peer error: %v\n", err)
	}
	for k, v := range message {
		fmt.Printf("%s : %q\n", k, v)
	}
	fmt.Printf("\n")
}
