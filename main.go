package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello world")
    })

    log.Println("Serving on localhost:3003")
    err := http.ListenAndServe(":3003", nil)
    log.Fatal(err)
}
