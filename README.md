goami
=====
Asterisk Manager Interface (AMI) client in Go.

## About
This code is based on C [libami](http://sourceforge.net/projects/amsuite/files/libami/) library interface

## Features

- Full AMI protocol implementation for Asterisk PBX
- Support for all major AMI actions (Login, Logoff, Originate, Queue operations, etc.)
- Event handling with automatic event subscription
- Concurrent connection management with context support
- Type-safe action parameters using Go structs with AMI tags
- Built-in support for:
  - SIP peer management
  - Queue operations
  - Configuration file management
  - AOC (Advice of Charge) data handling
  - Monitor/MixMonitor operations
  - Jabber messaging
  - Extension management
  - Message sending
- Custom action support for extending functionality
- Comprehensive error handling and response parsing
- Mock client support for testing

## Project Structure

```
goami/
├── ami/                    # Core AMI library package
│   ├── agi.go             # AGI control types
│   ├── client.go          # Client interface definition
│   ├── encode.go          # AMI message encoding/marshaling
│   ├── manager.go         # High-level AMI action functions
│   ├── monitor.go         # Monitor/MixMonitor data structures
│   ├── response.go        # Response type and parsing
│   ├── sip.go             # SIP-related AMI actions
│   ├── socket.go          # TCP socket connection management
│   ├── types.go           # Data structures for AMI actions
│   └── utils.go           # Utility functions (command building, sending)
├── example/               # Example implementation
│   └── asterisk.go        # High-level Asterisk client wrapper
├── README.md              # This file
└── LICENSE                # MIT License
```

The library is organized into two main packages:
- **`ami`**: The core library providing low-level AMI protocol implementation
- **`example`**: A reference implementation showing how to use the library

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

## Documentation

This projects documentation can be found on godoc at [goami](http://godoc.org/github.com/heltonmarx/goami/ami)
and supports:
 - *master*: [Asterisk 14 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+14+AMI+Actions)
 - ami.v10: [Asterisk 10 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+10+AMI+Actions)
 - ami.v13: [Asterisk 13 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+13+AMI+Actions)
 - ami.v14: [Asterisk 14 AMI Actions](https://wiki.asterisk.org/wiki/display/AST/Asterisk+14+AMI+Actions)

## Contributing

Contributions are welcome! Here's how you can help:

1. **Report bugs**: Open an issue describing the problem, including:
   - Go version
   - Asterisk version
   - Steps to reproduce
   - Expected vs actual behavior

2. **Submit fixes**: Create a pull request with:
   - Clear description of the problem and solution
   - Tests for new functionality
   - Updated documentation if needed
   - Code following Go conventions (`gofmt` compliant)

3. **Add new AMI actions**: Follow the existing pattern:
   - Define data structures in `types.go` with proper AMI tags
   - Add action functions in `manager.go`
   - Add tests in `manager_test.go`

4. **Improve documentation**: Update README, add examples, or improve code comments

### Development Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/heltonmarx/goami.git
   cd goami
   ```

2. Run tests:
   ```sh
   go test ./ami/...
   ```

3. Run tests with coverage:
   ```sh
   go test -cover ./ami/...
   ```

### Code Style

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable names
- Add comments for exported functions and types
- Write tests for new functionality
- Keep backward compatibility when possible

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/heltonmarx/goami/blob/master/LICENSE) file for details.

The MIT License is a permissive license that allows you to:
- Use the code commercially
- Modify the code
- Distribute the code
- Use it privately
- Sublicense the code

The only requirement is that you include the original copyright notice and license text in any copy of the software/source.

## Support

For questions, issues, or feature requests:
- Open an issue on [GitHub](https://github.com/heltonmarx/goami/issues)
- Check the [Asterisk AMI Documentation](https://wiki.asterisk.org/wiki/display/AST/Asterisk+Manager+Interface+AMI)
- Review existing issues and discussions

## Acknowledgments

- Based on the original C [libami](http://sourceforge.net/projects/amsuite/files/libami/) library
- Thanks to all contributors who have helped improve this library
- Special thanks to the Asterisk community for maintaining excellent documentation
