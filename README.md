goami
=====
Asterisk Manager Interface (AMI) client in Go 

## About
This code is based on C [libami] (http://sourceforge.net/projects/amsuite/files/libami/) library interface

## Installation and Requirements

The following command will install the notification api for Go along with the binaries.

```sh
go get github.com/heltonmarx/goami/ami
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
	"fmt"
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
	var ret bool

	//Login
	uuid, _ := ami.GetUUID()
	ret, err = ami.Login(socket, "admin", "admin", "Off", uuid)
	if err != nil || ret == false {
		fmt.Printf("login error (%v)\n", err)
		return
	}
	fmt.Printf("login ok!\n")

	//Logoff
	fmt.Printf("logoff\n")
	ret, err = ami.Logoff(socket, uuid)
	if err != nil || ret == false {
		fmt.Printf("logoff error: (%v)\n", err)
		return
	}
	fmt.Printf("goodbye !\n")
}
```

## Copyright and licensing

Distributed under GNU Lesser General Public License.
