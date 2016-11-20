package main

import (
	"fmt"
	"log"

	"github.com/heltonmarx/goami/ami"
)

func main() {
	socket, err := ami.NewSocket("127.0.0.1:5038")
	if err != nil {
		log.Fatalf("socket error: %v\n", err)
	}
	if _, err := ami.Connect(socket); err != nil {
		return
	}
	//Login
	uuid, _ := ami.GetUUID()
	if err := ami.Login(socket, "admin", "admin", "Off", uuid); err != nil {
		log.Fatalf("login error (%v)\n", err)
	}
	fmt.Printf("login ok!\n")

	//Logoff
	fmt.Printf("logoff\n")
	if err := ami.Logoff(socket, uuid); err != nil {
		log.Fatalf("logoff error: (%v)\n", err)
	}
	fmt.Printf("goodbye !\n")
}
