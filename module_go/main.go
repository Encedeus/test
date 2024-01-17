/*package main

import (
    "flag"
    "fmt"
    "os"

    // _ "github.com/stealthrocket/net/http"
    "github.com/stealthrocket/net/wasip1"
    "log"
    "net/http"
    "time"
)

type Handle struct{}

func main() {
    port := flag.Uint("p", 0, "Enter a port")
    flag.Parse()
    if *port == 0 {
        log.Fatalf("No valid port provided: %v", *port)
    }

    listener, err := wasip1.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
    if err != nil {
        log.Fatal(err)
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        writer.WriteHeader(200)
        writer.Header().Set("Content-Type", "text/html")
        f, _ := os.ReadFile("./public/index.html")
        writer.Write(f)
    })

    server := &http.Server{
        Addr:         fmt.Sprintf(":%v", *port),
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 5 * time.Second,
    }

    if err := server.Serve(listener); err != nil {
        log.Fatal(err)
    }

}*/

package main

import (
    "fmt"
    "github.com/labstack/echo/v4"
    _ "github.com/stealthrocket/net/http"
    "github.com/stealthrocket/net/wasip1"
    "log"
    "os"
    "strconv"
)

func main() {
    /*    port := flag.Uint("p", 0, "Enter a port")
          flag.Parse()
          if *port == 0 || *port > math.MaxUint16 {
              log.Fatalf("No valid port provided: %v", *port)
          }*/
    port, _ := strconv.Atoi(os.Getenv("MODULE_PORT"))

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        dir, _ := os.ReadDir("./")
        for _, i := range dir {
            fmt.Println(i.Name())
        }
        f, err := os.ReadFile("public/index.html")
        if err != nil {
            fmt.Printf("%e", err)
        }
        fmt.Println(string(f))
        return c.HTML(200, string(f))
    })

    listener, err := wasip1.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", port))
    if err != nil {
        log.Fatal(err)
    }
    e.Listener = listener

    e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}

/*package main

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
*/
