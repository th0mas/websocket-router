package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	log.Printf("Starting web server on %s", *addr)
	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		serveWs(rw, r)
	})

	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
