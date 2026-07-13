# AMI Connection Pool

This package provides a connection pool for managing multiple concurrent AMI
(Asterisk Manager Interface) connections to one or more Asterisk PBX servers.

## Features

- **Multi-server**: Connect to several Asterisk servers at once, each with its
  own credentials and event mask.
- **Automatic reconnect**: Each connection is kept alive with exponential
  backoff (1s → 30s max) when the remote end drops.
- **Single event channel**: Events from every server are multiplexed onto one
  channel, tagged with the server name so a single consumer can tell them apart.
- **Reference-counted sockets**: Callers that obtain a socket via `Get` or
  `Next` must call `Release` when done. The pool waits for all outstanding
  references before closing a socket on shutdown.
- **Round-robin selection**: `Next` distributes actions across healthy servers
  using an atomic counter.
- **Health snapshot**: `Healthy` returns a point-in-time map of which servers
  are currently connected.

## Usage

```go
servers := []ServerConfig{
    {Name: "sp-01", Host: "10.0.1.10:5038", Username: "admin", Secret: "admin"},
    {Name: "sp-02", Host: "10.0.2.10:5038", Username: "admin", Secret: "admin"},
}

ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

p := NewPool(servers)
p.Start(ctx)
defer p.Close()

// Consume events from every server.
go func() {
    for evt := range p.Events() {
        log.Printf("[%s] %s", evt.Server, evt.Data.Get("Event"))
    }
}()

// Target a specific server.
socket, actionID, err := p.Get("sp-01")
if err == nil {
    resp, _ := ami.CoreStatus(ctx, socket, actionID)
    log.Printf("sp-01: %s", resp.Get("CoreStartupDate"))
    p.Release("sp-01")
}

// Any healthy server.
name, socket, actionID, err := p.Next()
if err == nil {
    resp, _ := ami.CoreSettings(ctx, socket, actionID)
    log.Printf("%s: AMI version %s", name, resp.Get("AMIversion"))
    p.Release(name)
}
```

## Reference counting

Every successful `Get` or `Next` call increments an internal reference count
for that server's connection. The caller **must** call `Release(name)` when it
is done with the socket. Failing to do so will prevent the connection from
being closed cleanly on shutdown.

Calling `Release` twice for the same name (double-release) will cause a panic
because the underlying `sync.WaitGroup` counter becomes negative.

## Health

`Healthy()` returns a map of server names to booleans. It is a point-in-time
snapshot; the actual health of a server may change immediately after the call
returns. This is suitable for monitoring endpoints or metrics gauges, not for
making decisions about individual actions (use `Get` or `Next` for that).
