# Logd
![A circular buffer](.doc/circular_buffer.svg)
## Logs for your apps in constant time and constant space with ultra-low latency.
Logd is a circular buffer for writing & reading millions of logs per minute.


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
Logd is built on Protobuf. Messages are serialised using the `msg.Msg` type generated from the Protobuf definition.

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
See the following example. Error checks skipped for brevity.
```go
// dial udp
addr, _ := conn.GetAddr("logd.fly.dev")
socket, _ := conn.Dial(addr)

// serialise message using protobuf
payload, _ := proto.Marshal(&msg.Msg{
  Timestamp: time.Now().UnixNano(),
  Env:       l.Env,
  Svc:       l.Svc,
  Fn:        l.Fn,
  Lvl:       string(lvl),
  Msg:       fmt.Sprintf(template, args...),
})

// get ephemeral signature using current time
signedMsg, _ := auth.Sign("some-secret-value", payload, time.Now())

// write to socket
socket.Write(signedMsg)
```
