package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", Greeting)
    http.ListenAndServe(":80", nil)
}

func Greeting(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<b>Hello, %s! </b>", r.URL.Path[1:])
}