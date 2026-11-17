package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type Product struct {
  ID int          `json:"id"`
  Name string     `json:"name"`
  Price float64   `json:"price"`
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    product := &Product{
      ID: 1,
      Name: "Car",
      Price: 200000,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}


func CreateProduct(w http.ResponseWriter, r *http.Request)  {
    var product Product

   err := json.NewDecoder(r.Body).Decode(&product)

   if err != nil {
      http.Error(w, "Invalid Json" ,http.StatusBadRequest)

      return
   }

   if product.Name == "" {
    http.Error(w, "Name is required", http.StatusBadRequest)

    return
   }

   if product.Price <= 0 {
    http.Error(w, "Price must be greater than 0", http.StatusBadRequest)

    return
   }

   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusCreated)
   json.NewEncoder(w).Encode(product)
}

func GetProductDetail(w http.ResponseWriter, r *http.Request)  {
  id := r.PathValue("id")
  fmt.Fprintln(w, id)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request)  {
  id := r.PathValue("id")
  fmt.Fprintln(w, id)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request)  {
  id := r.PathValue("id")
  fmt.Fprintln(w, id)
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Request:", r.Method, r.URL.Path)

        next.ServeHTTP(w, r)
    })
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")

        if token != "Bearer secret123" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {

    mux := http.NewServeMux()

    mux.HandleFunc("GET /products", GetProduct)
    mux.HandleFunc("GET /products/{id}", GetProductDetail)

    mux.Handle("POST /products", AuthMiddleware(http.HandlerFunc(CreateProduct)))
    mux.Handle("PUT /products/{id}", AuthMiddleware(http.HandlerFunc(UpdateProduct)))
    mux.Handle("DELETE /products/{id}", AuthMiddleware(http.HandlerFunc(DeleteProduct)))

    handler := CORSMiddleware(LoggingMiddleware(mux))

    http.ListenAndServe(":8080", handler)
}