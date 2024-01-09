# Logd
![A circular buffer](.doc/circular_buffer.svg)
## Logs for your apps in constant time and constant space with ultra-low latency.
Logd is a non-blocking circular buffer to store millions of logs per hour.


# Logd is a circular buffer
Logd (pronounced "logged") will never run out of storage. Reads & writes are constant time.

As the buffer becomes full, each write overwrites the oldest element.

# Auth
Logd authenticates clients for either reading or writing using 2 shared secrets.
These are stored encrypted in our secrets SOPS file.

# HTTP API
Logd starts a http server.
### GET /
```bash
curl --location "$LOGD_HOST/?limit=10" \
--header "Authorization: $LOGD_READ_SECRET"
```

# UDP
## Logger
The simplest way to write logs is using the `log` package.
```go
l, _ := log.NewLogger(&log.LoggerConfig{
  Host:        "logd.fly.dev",
  WriteSecret: "the-secret",
  Env:         "prod",
  Svc:         "example-service",
  Fn:          "Readme",
})
l.Log(log.Info, "this is an example %s", "log message")
```

## Custom integration
Logs are written by connecting to a UDP socket on port `:6102`.
```go
// error checks skipped for brevity

// dial udp
addr, _ := conn.GetAddr("logd.fly.dev")
conn, _ := conn.Dial(addr)

// serialise message
payload, _ := pack.PackMsg(&msg.Msg{
  Timestamp: time.Now().UnixNano(),
  Env:       l.Env,
  Svc:       l.Svc,
  Fn:        l.Fn,
  Lvl:       string(lvl),
  Msg:       fmt.Sprintf(template, args...),
})

// get ephemeral signature using current time
signedMsg, _ := auth.Sign(l.Secret, payload, time.Now())

// write to socket
l.Conn.Write(signedMsg)
```
