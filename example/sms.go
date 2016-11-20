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

	data := ami.KhompSMSData{
		Device:       "b0",
		Destination:  "4899893791",
		Confirmation: true,
		Message:      "hey ho, let's go",
	}
	s, err := ami.KSendSMS(socket, uuid, data)
	if err != nil {
		fmt.Printf("sms sms error\n", err)
	}
	fmt.Printf("response: [%v]\n", s)
}
