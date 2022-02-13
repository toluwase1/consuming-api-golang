package main

import (
	"log"
	"net/http"
	"time"
)

var client2  *http.Client


func main() {
	client2 = &http.Client{Timeout: 10 * time.Second}
	http.HandleFunc("/hello", simpleServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func simpleServer (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My Simple Web Server"))
}