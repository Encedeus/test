package main

import (
    "log"
    "net/http"
    "time"

    "github.com/stealthrocket/net/wasip1"
)

type TestHandler struct{}

func (th TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func main() {
    listener, err := wasip1.Listen("tcp", "127.0.0.1:6969")
    if err != nil {
        log.Fatal(err)
    }

    // mux := http.NewServeMux()
    // mux.HandleFunc("hello", func(w http.ResponseWriter, r *http.Request) {
    //     w.Write([]byte("Hello, world!"))
    //     w.WriteHeader(http.StatusOK)
    // })

    server := &http.Server{
        Addr:         ":6969",
        Handler:      TestHandler{},
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    if err := server.Serve(listener); err != nil {
        log.Fatal(err)
    }
}
