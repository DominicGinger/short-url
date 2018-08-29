package main

import (
    "math/rand"
    "fmt"
    "log"
    "net/http"
    "time"
)

const port = "3003"
const storageLimit = 62*62

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    data := make(map[string]string)

    http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
        params := r.URL.Query()
        key := params.Get("key")
        value := params.Get("value")

        if (key == "") {
            key = randomKey()
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

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Path[1:]

        if (key == "") {
            fmt.Fprintln(w, "Missing key param, expected ?key=x")
            return
        }

        http.Redirect(w, r, data[key], http.StatusSeeOther)
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

func randomKey() string {
    letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    b := make([]rune, 2)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}

