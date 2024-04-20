package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}

func main() {
    http.HandleFunc("/items", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
