package main

import (
    "github.com/filecoin-project/go-jsonrpc"
    "net/http"
    "time"
)

type TestHandler struct {
    n int
}

func (h *TestHandler) AddGet(in int) int {
    h.n += in
    return h.n
}

func main() {
    rpcServer := jsonrpc.NewServer()

    sHandler := new(TestHandler)
    rpcServer.Register("TestHandler", sHandler)

    s := http.Server{
        Handler:      rpcServer,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
        Addr:         ":8082",
    }

    s.ListenAndServe()
}
