package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	asterisk, err := NewAsterisk(ctx, *host, *username, *secret)
	if err != nil {
		log.Fatal(err)
	}
	defer asterisk.Logoff(ctx)

	log.Printf("connected with asterisk\n")

	peers, err := asterisk.SIPPeers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("peers: %v\n", peers)
}
