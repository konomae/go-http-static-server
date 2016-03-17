package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	port := flag.Int("port", 8080, "port number")
	host := flag.String("host", "127.0.0.1", "host")
	path := flag.String("path", ".", "document root path")
	keyPath := flag.String("key", "", "private key path")
	pubPath := flag.String("pub", "", "public key path")

	flag.Parse()

	absPath, _ := filepath.Abs(*path)
	hostPort := fmt.Sprintf("%s:%d", *host, *port)
	handler := logger(http.FileServer(http.Dir(absPath)))

	log.Println("Path:", absPath)

	if *keyPath == "" {
		log.Println("URL :", "http://"+hostPort)
		log.Fatal(http.ListenAndServe(hostPort, handler))
	} else {
		log.Println("URL :", "https://"+hostPort)
		log.Fatal(http.ListenAndServeTLS(hostPort, *pubPath, *keyPath, handler))
	}

}
