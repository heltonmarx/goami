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
	var answer string
	var ret bool

	fmt.Printf("login\n")
	uuid, _ := ami.GetUUID()
	answer, err = ami.Login(socket, "admin", "admin", "Off", uuid)
	if err != nil {
		fmt.Printf("login error\n")
	}
	fmt.Printf("answer[%s]\n", answer)

	fmt.Printf("ping\n")
	ret, err = ami.Ping(socket, uuid)
	if err != nil {
		fmt.Printf("ping error\n")
	}
	fmt.Printf("[ok] - [%d]\n",ret)

	fmt.Printf("logoff\n")
	answer, err = ami.Logoff(socket, uuid)
	if err != nil {
		fmt.Printf("logoff error\n")
	}
	fmt.Printf("answer[%s]\n", answer)
}
