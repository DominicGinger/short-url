package main

import (
    "crypto/rand"
    "fmt"
    "log"
    "net/http"
)

const port = "3003"
const storageLimit = 5000

func main() {
    data := make(map[string]string)

    http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
        params := r.URL.Query()
        key := params.Get("key")
        value := params.Get("value")

        if (key == "") {
            key = randomKey(2)
        }

        if (value == "") {
            fmt.Fprintln(w, "Missing value param, expected ?value=y")
            return
        }
        if (!validKey(key)) {
            fmt.Fprintln(w, "Key too long, must be less than 16 chars")
            return
        }
        if (!validValue(value)) {
            fmt.Fprintln(w, "Value too long, must be less than 1000 chars")
            return
        }

        if (len(data) >= storageLimit) {
            data = make(map[string]string)
        }

        data[key] = value

        fmt.Fprintln(w, key)
    })

    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
        params := r.URL.Query()
        key := params.Get("key")

        if (key == "") {
            fmt.Fprintln(w, "Missing key param, expected ?key=x")
            return
        }

        fmt.Fprintln(w, data[key])
    })

    log.Println("Listening on " + port)
    err := http.ListenAndServe(":" + port, nil)
    log.Fatal(err)
}

func validKey(key string) bool {
    return len(key) <= 16
}

func validValue(value string) bool {
    return len(value) <= 1000
}

func randomKey(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        panic(err)
    }
    return fmt.Sprintf("%X", b)
}
