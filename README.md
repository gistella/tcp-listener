# tcp-listener

Simple go implementation of a TCP listener.
Source connections' IP/port are logged and printed back on STDOUT.

By default the listener listens on port 8080 but another port can be set via the `LISTEN_PORT` env variable.

## Compiling and running

```
$ go build -o tcplisten main.go
$ ./tcplisten
Listening on port 8080
^C

# listen on another port
$ LISTEN_PORT=10142 ./tcplisten
Listening on port 10142
^C
```
