# go-tcp-echo

> A one-off tcp echo server

```
Usage of ./go-tcp-echo:
  -port int
    	Port to accept connections on. (default 7777)

./go-tcp-echo 
2017/02/16 10:19:08 Listening to connections on port 7777


# then, from another terminal

telnet localhost 7777
Trying ::1...
Connected to localhost.
Escape character is '^]'.
example
example
Connection closed by foreign host.
```

## Docker image

Just pull `cirocosta/go-tcp-echo` and run it:

```
docker run -d -p "7777:7777" cirocosta/go-tcp-echo 
```


## Developing

See `./Makefile`.

