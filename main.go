package main

import (
	"fmt"
	"net/http"
)

func UserRequest(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")

    switch r.Method {
    case http.MethodGet:
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Getting users")

    case http.MethodPost:
        w.WriteHeader(http.StatusCreated)
        fmt.Fprintln(w, "Creating user")

    case http.MethodDelete:
        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(
            w,
            "Method not allowed",
            http.StatusMethodNotAllowed,
        )
    }
}

func main() {
    http.HandleFunc("/users", UserRequest)

    http.ListenAndServe(":8080", nil)
}