package main

import (
	"fmt"
	"net/http"
)

func ProductRequest(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "Product API is running")
}

func main() {
    http.HandleFunc("/products", ProductRequest)

    http.ListenAndServe(":8080", nil)
}