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

## Connections Pool
For concurrent (i.e. REST API) connections you may need keep a pool of AMI connections
```Go

	events := "system,call,all,user"
	pool, err := ami.NewPool(ctx, host, username, secret, events)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating pool")
		return
	}
	defer pool.CloseAll()
	pool.MinConections = 2 	//minimun ami sessions to  keep alive
	pool.MaxConections = 20 //max allowed concurrent sessions to AMI


	//this will start two sockets to asterisk because `pool.LowWater = 2`
	if err := pool.Connect(); err != nil {
		log.Fatal().Err(err).Msg("cant connect with asterisk")
	}

	//get an socket from the pool
	s1, err := pool.GetSocket() //get a socket
	s2, err := pool.GetSocket() //get another socket
	s3, err := pool.GetSocket() //starts new socket to asterisk ...
	defer pool.Close(s1, false) //don't forget to give back the connection to the pool!!! 
	defer pool.Close(s2, false) //don't forget to give back the connection to the pool!!! 
	defer pool.Close(s3, false) //don't forget to give back the connection to the pool!!! 
	
	if err := ami.Ping(ctx, s1, ""); err != nil {
		//something went wrong with this connection. kill it!
		pool.Close(s1, true)
	}
	ami.Ping(ctx, s2, "");
	ami.Ping(ctx, s3, "");
```
### Getting Error: "Max allowed connections reached. Increase the max allowed connections by setting pool.MaxConections to a higher value"
If you get this error with even low traffic, ensure you "gived back" the socket to the pool with 
```Go
defer pool.Close(<socket>, false)
```




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
