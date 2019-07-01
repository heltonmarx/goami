goami
=====
Asterisk Manager Interface (AMI) client in Go.

## About
This code is based on C [libami](http://sourceforge.net/projects/amsuite/files/libami/) library interface

## Installation and Requirements

The following command will install the AMI client.

```sh
go get -u github.com/heltonmarx/goami/ami
```

To test this package with Asterisk it's necessary set the file `/etc/asterisk/manager.conf` with configuration bellow:

    [general]
    enabled = yes
    port = 5038
    bindaddr = 127.0.0.1

    [admin]
    secret = admin
    deny = 0.0.0.0/0.0.0.0
    permit = 127.0.0.1/255.255.255.255
    read = all,system,call,log,verbose,command,agent,user,config
    write = all,system,call,log,verbose,command,agent,user,config

## Using the code

Login/Logoff:
```Go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/heltonmarx/goami/ami"
)

var (
	username = flag.String("username", "admin", "AMI username")
	secret   = flag.String("secret", "admin", "AMI secret")
	host     = flag.String("host", "127.0.0.1:5038", "AMI host address")
)

func main() {
	flag.Parse()

	socket, err := ami.NewSocket(*host)
	if err != nil {
		log.Fatalf("socket error: %v\n", err)
	}
	if _, err := ami.Connect(socket); err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	//Login
	uuid, _ := ami.GetUUID()
	if err := ami.Login(socket, *username, *secret, "Off", uuid); err != nil {
		log.Fatalf("login error: %v\n", err)
	}
	fmt.Printf("login ok!\n")

	//Logoff
	fmt.Printf("logoff\n")
	if err := ami.Logoff(socket, uuid); err != nil {
		log.Fatalf("logoff error: (%v)\n", err)
	}
	fmt.Printf("goodbye !\n")
}
```

## Documentation

This projects documentation can be found on godoc at [goami](http://godoc.org/github.com/heltonmarx/goami/ami),
and supports [Asterisk 10 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+10+AMI+Actions)

## License

MIT-LICENSE. See [LICENSE](https://github.com/heltonmarx/goami/blob/master/LICENSE)
or the LICENSE file provied in the repository for details.

