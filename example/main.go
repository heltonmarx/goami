package main

import (
	"flag"
	"log"
)

var (
	username = flag.String("username", "admin", "AMI username")
	secret   = flag.String("secret", "dials", "AMI secret")
	host     = flag.String("host", "127.0.0.1:5038", "AMI host address")
)

func main() {
	flag.Parse()

	asterisk, err := NewAsterisk(*host, *username, *secret)
	if err != nil {
		log.Fatal(err)
	}
	defer asterisk.Logoff()

	peers, err := asterisk.SIPPeers()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("peers: %v\n", peers)
}
