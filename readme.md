# go-http-static-server

```bash
$ go-http-static-server -h

Usage of go-http-static-server:
  -host string
    	host (default "127.0.0.1")
  -key string
    	private key path
  -path string
    	document root path (default ".")
  -port int
    	port number (default 8080)
  -pub string
    	public key path
```

## Install

```bash
$ go get -u github.com/konomae/go-http-static-server
```


## Generate a certificate

```bash
$ openssl req -new -newkey rsa:2048 -x509 -keyout server.key -out server.pem -days 365 -nodes
```
