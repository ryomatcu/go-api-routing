# go-api-routing
Sample API routing with Golang

# Usage

Install package

```
$ go get "github.com/julienschmidt/httprouter"
```

Run main.go

```
$ go run main.go
```

Request localhost

```
$ curl -vvv http://localhost:8080
* Rebuilt URL to: http://localhost:8080/
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 26 Dec 2017 11:58:36 GMT
< Content-Length: 8
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Welcome!
```

```
$ curl http://localhost:8080/items
Item Index!

$ curl http://localhost:8080/items/1
Item show: 1
```