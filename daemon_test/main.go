package main

import (
    "fmt"
    daemon "github.com/Encedeus/module-daemon-go"
    "github.com/Encedeus/module-daemon-go/module"
    "github.com/labstack/echo/v4"
    "github.com/stealthrocket/net/wasip1"
    "log"
)

func main() {
    fmt.Println("Hello, world mate!")
    daemon.InitModule(func(m *module.Module) {
        fmt.Println("Lol")
        m.RegisterCommand(module.Command{
            Name:   "test_cmd",
            Params: nil,
            Exec: func(m *module.Module, args module.Arguments) (module.Result, error) {
                go func() {
                    fmt.Println("Hello, world!")
                    e := echo.New()
                    e.GET("/", func(c echo.Context) error {
                        return c.String(200, "Hello, world!")
                    })

                    listener, err := wasip1.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", 8086))
                    if err != nil {
                        log.Fatal(err)
                    }
                    e.Listener = listener

                    e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", 8086)))
                }()

                return nil, nil
            },
        })

        // time.Sleep(2 * time.Second)
        _, err := m.Invoke("test_cmd", nil)
        if err != nil {
            log.Printf("Error: %e", err)
        }
    })
}
