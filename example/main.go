package main

import (
	"flag"
	"log"
)

var (
	username = flag.String("username", "admin", "AMI username")
	secret   = flag.String("secret", "admin", "AMI secret")
	host     = flag.String("host", "127.0.0.1:5038", "AMI host address")
)

func main() {
	flag.Parse()

	asterisk, err := NewAsterisk(*host, *username, *secret)
	if err != nil {
		log.Fatal(err)
	}
	defer asterisk.Logoff()

	err = asterisk.SIPPeers()
	if err != nil {
		log.Fatal(err)
	}
}
