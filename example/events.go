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
	if err := ami.Login(socket, "admin", "admin", "system,call,all,user", uuid); err != nil {
		log.Fatalf("login error (%v)\n", err)
	}
	defer ami.Logoff(socket, uuid)
	fmt.Printf("login ok!\n")

	//Events
	for {
		events, err := ami.Events(socket)
		if err != nil {
			fmt.Printf("events error (%v)\n", err)
			return
		}
		fmt.Printf("recv event: %v\n", events)
	}
}
