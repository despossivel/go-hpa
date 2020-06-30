package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", Greeting)
    http.ListenAndServe(":8080", nil)
}

func Greeting(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}