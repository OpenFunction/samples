package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/sample-topic", func(rw http.ResponseWriter, req *http.Request) {
        var msg Message

        err := json.NewDecoder(req.Body).Decode(&msg)
        if err != nil {
            fmt.Println("error reading message from Kafka binding", err)
            rw.WriteHeader(500)
            return
        }
        fmt.Printf("message from Kafka '%s'\n", msg)
        rw.WriteHeader(200)
    })
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
}

type Message struct {
    Msg string `json:"message"`
}
