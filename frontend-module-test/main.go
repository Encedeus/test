package main

import (
    vueglue "github.com/torenware/vite-go"
    "html/template"
    "io/fs"
    "log"
    "mime"
    "net/http"
    "os"
    "path/filepath"
)

var dist fs.FS

var vueGlue *vueglue.VueGlue

func main() {
    config := &vueglue.ViteConfig{
        Environment: "production",
        EntryPoint:  "src/main.ts",
        Platform:    "react",
    }

    dist = os.DirFS("frontend-src-react")
    config.FS = dist

    glue, err := vueglue.NewVueGlue(config)
    if err != nil {
        panic(err)
    }
    vueGlue = glue

    mux := http.NewServeMux()

    fsHandler, err := glue.FileServer()
    if err != nil {
        panic(err)
    }

    mux.Handle(config.URLPrefix, fsHandler)
    mux.Handle("/", http.HandlerFunc(pageWithAVue))

    http.ListenAndServe(":4000", mux)
}

func serveOneFile(w http.ResponseWriter, r *http.Request, uri, contentType string) {
    strippedURI := uri[1:]
    buf, err := os.ReadFile("frontend-src-react/dist/" + strippedURI)

    if err != nil {
        w.WriteHeader(http.StatusNotFound)
    }

    w.Header().Add("Content-Type", contentType)
    w.Write(buf)
}

func pageWithAVue(w http.ResponseWriter, r *http.Request) {

    ext := filepath.Ext(r.RequestURI)
    /*    re := regexp.MustCompile(`^/([^.]+)\.(svg|ico|jpg|png)$`)
          matches := re.FindStringSubmatch(r.RequestURI)
          if matches != nil {
              var contentType string

              switch matches[2] {
              case "svg":
                  contentType = "image/svg+xml"
              case "ico":
                  contentType = "image/x-icon"
              case "jpg":
                  contentType = "image/jpeg"
              case "png":
                  contentType = "image/png"
              }

              serveOneFile(w, r, r.RequestURI, contentType)
              return
          }*/

    contentType = mime.TypeByExtension(ext)
    t, err := template.ParseFiles("./test-template.gohtml")
    if err != nil {
        log.Fatal(err)
    }

    t.Execute(w, vueGlue)
}
