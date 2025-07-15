# myip

An HTTP server that responds with client's IP.

## Build and run

```plain
$ go build -o ./bin/myip
$ ./bin/myip -port 8080
2025/07/15 16:15:47 Listening on :8080...

$ curl http://192.168.1.1:8080
192.168.1.10
```
