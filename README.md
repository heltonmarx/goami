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
├── example/               # Example implementations
│   ├── simple/            # Single-server example (see example/simple/)
│   │   ├── asterisk.go    # High-level Asterisk client wrapper
│   │   ├── main.go        # Entry point for the simple example
│   │   ├── Makefile       # Build/test targets
│   │   └── docker-compose.yml  # Asterisk container for local testing
│   └── pool/              # Multi-server connection pool (see example/pool/README.md)
│       ├── main.go        # Entry point for the pool example
│       ├── pool.go        # Pool implementation
│       ├── Makefile       # Build/test targets
│       └── README.md      # Pool-specific documentation
├── README.md              # This file
└── LICENSE                # MIT License
```

The library is organized into two main packages:
- **`ami`**: The core library providing low-level AMI protocol implementation
- **`example`**: Reference implementations showing how to use the library

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

## Examples

### Simple (single-server) example

The `example/simple/` directory contains a complete single-server AMI client that connects to one Asterisk instance, logs in, fetches SIP peers, and prints them.

**Code** (`example/simple/main.go`):

```go
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
```

**Run instructions:**

1. Start an Asterisk container (requires Docker):
   ```sh
   cd example/simple
   docker-compose up -d
   ```
2. Build and run the example:
   ```sh
   cd example/simple
   go build -o simple .
   ./simple -host 127.0.0.1:5038 -username admin -secret dials
   ```

The example uses the `docker-compose.yml` file in `example/simple/` to spin up a local Asterisk instance for testing.

### Pool (multi-server) example

The `example/pool/` directory contains a connection pool that manages multiple concurrent AMI connections to several Asterisk servers, with automatic reconnect and round‑robin selection.

**Code** (`example/pool/main.go`):

```go
// Command poolexample connects to several Asterisk servers concurrently
// using the pool package, prints incoming events tagged by server, and
// shows how to target a single server for an action-style call.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/heltonmarx/goami/ami"
)

func main() {
	username := flag.String("username", "admin", "AMI username (shared by all servers in this example)")
	secret := flag.String("secret", "admin", "AMI secret (shared by all servers in this example)")
	flag.Parse()

	// In a real service these would come from config/env/service discovery,
	// each with its own credentials if needed.
	servers := []ServerConfig{
		{Name: "sp-01", Host: "10.0.1.10:5038", Username: *username, Secret: *secret},
		{Name: "sp-02", Host: "10.0.2.10:5038", Username: *username, Secret: *secret},
		{Name: "sp-03", Host: "10.0.3.10:5038", Username: *username, Secret: *secret},
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	p := NewPool(servers)
	p.Start(ctx)
	defer p.Close()

	// Consume events from every server on one channel.
	go func() {
		for evt := range p.Events() {
			log.Printf("[%s] event=%s channel=%s",
				evt.Server, evt.Data.Get("Event"), evt.Data.Get("Channel"))
		}
	}()

	// Example: target one specific server for an action.
	if socket, actionID, err := p.Get("sp-02"); err == nil {
		if resp, err := ami.CoreStatus(ctx, socket, actionID); err == nil {
			log.Printf("sp-02 core status: %s", resp.Get("CoreStartupDate"))
		}
		p.Release("sp-02")
	}

	// Example: don't care which server, just need any healthy one.
	if name, socket, actionID, err := p.Next(); err == nil {
		if resp, err := ami.CoreSettings(ctx, socket, actionID); err == nil {
			log.Printf("picked %s via round-robin, AMI version %s", name, resp.Get("AMIversion"))
		}
		p.Release(name)
	}

	<-ctx.Done()
	log.Println("shutting down")
}
```

**Run instructions:**

1. Ensure you have at least one Asterisk instance reachable (the example expects three servers; adjust the `servers` slice to match your environment).
2. Build and run:
   ```sh
   cd example/pool
   go build -o poolexample .
   ./poolexample -username admin -secret admin
   ```

For full details on the pool implementation, see `example/pool/README.md`.

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
