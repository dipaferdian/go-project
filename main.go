package main

import (
	"fmt"
	"net/http"
	"strconv"
)


func UserDetail(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    fmt.Fprintln(w, "User ID:", id)
}

func UserList(w http.ResponseWriter, r *http.Request) {
    page, err := strconv.Atoi(r.URL.Query().Get("page"))

    if err != nil {
      http.Error(w, "Invalid page", http.StatusBadRequest)
      return
    }

    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

    if err != nil {
      http.Error(w, "Invalid page", http.StatusBadRequest)
      return
    }

    fmt.Fprintln(w, "User page:", page, "User limit:", limit)
}


func main() {
    http.HandleFunc("GET /users/{id}", UserDetail)
    http.HandleFunc("GET /users", UserList)

    http.ListenAndServe(":8080", nil)
}