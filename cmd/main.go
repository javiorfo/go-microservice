package main

import (
	"fmt"
	"log"
	"net/http"
    "github.com/javiorfo/go-microservice/adapter/out"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}

func main() {
    db := out.GetConnection()
    var dummies []out.Dummy
    db.Find(&dummies)
    fmt.Println(dummies)
    http.HandleFunc("/items", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
