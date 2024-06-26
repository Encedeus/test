package main

import (
    "context"
    "fmt"
    "github.com/filecoin-project/go-jsonrpc"
)

func main() {
    _, err := Invoke("test_cmd", nil, 16959)
    fmt.Printf("Invoke error: %e", err)
}

type Result any
type Parameters []string
type Arguments map[string]any
type Executor func(m *Module, args Arguments) (Result, error)

type Command struct {
    Name   string
    Params Parameters
    Exec   Executor
}

type InvokeFunc func(command string, args Arguments) (Result, error)

type HostInvokeHandler struct {
    Module *Module
}

func (h *HostInvokeHandler) HostInvoke(command string, args Arguments) (Result, error) {
    for _, cmd := range h.Module.Commands {
        if cmd.Name == command {
            result, err := cmd.Exec(h.Module, args)
            if err != nil {
                return nil, err
            }

            return result, nil
        }
    }

    return nil, nil
}

type RunFunction func(m *Module)

type HandshakeHandler struct {
    RegisteredCommands []*Command
    Module             *Module
    Run                RunFunction
    RPCPort            Port
    MainPort           Port
}

type HandshakeResponse struct {
    // RegisteredCommands []*command.Command
}

func (h *HandshakeHandler) OnHandshake(config Configuration) HandshakeResponse {
    h.Module.Port = config.Port
    h.Module.Manifest = config.Manifest
    h.Module.HostPort = config.HostPort
    h.Module.HandshakeHandler = h

    defer h.Run(h.Module)
    fmt.Println("Handshake done from module side")

    return HandshakeResponse{}
}

type Manifest struct {
    Name             string   `hcl:"name"`
    Authors          []string `hcl:"authors"`
    Version          string   `hcl:"version"`
    FrontendMainFile string   `hcl:"frontend_main"`
    // BackendMainFile  string   `hcl:"backend_main"`
}

type Module struct {
    Port             Port
    Manifest         Manifest
    HostPort         Port
    Commands         []*Command
    HandshakeHandler *HandshakeHandler
}

func (m *Module) RegisterCommand(cmd Command) {
    m.Commands = append(m.Commands, &cmd)
}

func Invoke(cmd string, args Arguments, port Port) (Result, error) {
    var client struct {
        ModuleInvoke InvokeFunc
    }
    fmt.Printf("Host port: %v\n", port)
    // time.Sleep(2 * time.Second)

    /*    conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", m.HostPort))
          fmt.Printf("Invoke dial error: %e\n", err)
          if err != nil {
              log.Fatalf("%e", err)
          }
          defer conn.Close()*/

    closer, err := jsonrpc.NewClient(context.Background(), fmt.Sprintf("http://localhost:%v", port), "ModuleInvokeHandler", &client, nil)
    fmt.Printf("Client open error: %e\n", err)
    if err != nil {
        return nil, err
    }
    defer closer()

    result, err := client.ModuleInvoke(cmd, args)
    fmt.Printf("Module invoke error: %e\n", err)
    if err != nil {
        return nil, err
    }

    return result, nil
}

type Port uint16

type Configuration struct {
    Port     Port
    HostPort Port
    Manifest Manifest
}
