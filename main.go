package main

import (
    "fmt"
    "log"
    "net/http"
)

func newHandler() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        http.ServeFile(w, r, "index.html")
    })

    mux.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "styles.css")
    })

    return mux
}

func main() {
    fmt.Println("Portfolio website running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", newHandler()))
}
