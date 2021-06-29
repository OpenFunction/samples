package main

import (
	"io"
	"log"
	"net/http"
)

func echo(w http.ResponseWriter, req *http.Request) {
	content, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	log.Println("Receive a message:")
	log.Println(string(content))
}

func main() {

	http.HandleFunc("/echo", echo)

	err := http.ListenAndServe(":7489", nil)
	if err != nil {
		return
	}
}
