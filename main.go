package main

import (
    "fmt"
    "log"
    "net/http"
)

const port = "3003"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello world")
    })

    log.Println("Serving on " + port)
    err := http.ListenAndServe(":" + port, nil)
    log.Fatal(err)
}
