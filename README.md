goami
=====

AMI Interface in Go

## About
This code is based on C [libami] (http://sourceforge.net/projects/amsuite/files/libami/) library interface

## Usage instructions

To test this package with Asterisk it's necessary set the tile `/etc/asterisk/manager.conf` with configuration bellow:

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

## Copyright and licensing

Distributed under GNU Lesser General Public License.
