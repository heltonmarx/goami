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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	socket, err := ami.NewSocket(ctx, *host)
	if err != nil {
		log.Fatalf("socket error: %v\n", err)
	}
	if _, err := ami.Connect(ctx, socket); err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	//Login
	uuid, _ := ami.GetUUID()
	if err := ami.Login(ctx, socket, *username, *secret, "Off", uuid); err != nil {
		log.Fatalf("login error: %v\n", err)
	}
	fmt.Printf("login ok!\n")

	//Logoff
	fmt.Printf("logoff\n")
	if err := ami.Logoff(ctx, socket, uuid); err != nil {
		log.Fatalf("logoff error: (%v)\n", err)
	}
	fmt.Printf("goodbye !\n")
}
```


## Event Broker
by default, this module executes commands in synchronous mode (ie: send Command and wait Response). 
AMI connections are asynchronous.
You can enable asynchronous mode adding next code just before the Login Command

```Go
	// Start Event Broker
	socket.Events().Start(ctx, socket)
```
Now, you can subscribe for events:

```Go
	events := h.socket.Events().Subscribe()
	defer func() {
		h.socket.Events().Unsubscribe <- events
	}()

	for {
		select {
		//canceled
		case <-ctx.Done():
			return
		// received event from asterisk
		case response := <-events:
			js, _ := json.Marshal(response)
			fmt.Printf(w, "data: %s\n\n", js)
		}
	}

```
### Advantages of using Event Broker
* Multiple subscribers (ie: http request por events)
* The Action ID from response is compared width the action ID from command, so you have accurate response

## Empty actionID
If the actionID parameter is set to an empty string, a random uuid will be generated

## Documentation

This projects documentation can be found on godoc at [goami](http://godoc.org/github.com/heltonmarx/goami/ami)
and supports:
 - *master*: [Asterisk 14 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+14+AMI+Actions)
 - ami.v10: [Asterisk 10 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+10+AMI+Actions)
 - ami.v13: [Asterisk 13 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+13+AMI+Actions)
 - ami.v14: [Asterisk 14 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+14+AMI+Actions)

## License

MIT-LICENSE. See [LICENSE](https://github.com/heltonmarx/goami/blob/master/LICENSE)
or the LICENSE file provided in the repository for details.
