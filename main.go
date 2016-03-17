package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	// without program name
	args := os.Args[1:]
	hostPort := "127.0.0.1:8080"
	path := "."
	if len(args) >= 1 {
		path = args[0]
	}
	if len(args) >= 2 {
		hostPort = args[1]
	}

	absPath, _ := filepath.Abs(path)

	log.Println("Host:", hostPort)
	log.Println("Path:", absPath)
	log.Println("URL :", "http://" + hostPort)

	log.Fatal(http.ListenAndServe(hostPort, logger(http.FileServer(http.Dir(absPath)))))
}
