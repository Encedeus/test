package main

import (
    "context"
    "fmt"
    "github.com/filecoin-project/go-jsonrpc"
    "log"
)

func main() {

    var client struct {
        AddGet func(int) int
    }

    closer, err := jsonrpc.NewClient(context.Background(), "http://localhost:8082", "TestHandler", &client, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer closer()

    n := client.AddGet(69420)
    fmt.Printf("n: %v\n", n)
}
