package main

import (
    "context"
    "fmt"
    "log"

    protoapi "github.com/Encedeus/protobuf/panel"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial(":8080", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }

    p := protoapi.NewContainerClient(conn)

    create, err := p.Create(context.Background(), &protoapi.ContainerCreateParams{
        Config: &protoapi.ContainerConfig{
            Image:        "itzg/minecraft-server",
            Tty:          true,
            Env:          []string{"EULA=true"},
            AttachStdin:  true,
            AttachStderr: true,
            AttachStdout: true,
        },
        Name: "hello_world",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v", create)
}
