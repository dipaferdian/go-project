package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)


type Product struct {
  ID int          `json:"id"`
  Name string     `json:"name"`
  Price float64   `json:"price"`
}

type ErrorResponse struct {
    Error  string `json:"error"`
    Status int    `json:"status"`
}

var ErrProductNotFound = errors.New("product not found")
var ErrUnauthorized = errors.New("unauthorized")


func GetProduct(w http.ResponseWriter, r *http.Request) {
    product := &Product{
      ID: 1,
      Name: "Car",
      Price: 200000,
    }

     id, err := strconv.Atoi(r.URL.Query().Get("id"))

     if err != nil {
      writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request", Status: http.StatusBadRequest})
     }

     if id == 0 {
       writeJSON(w, http.StatusNotFound, ErrorResponse{Error: ErrProductNotFound.Error(), Status: http.StatusNotFound})
       return
     }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}


func CreateProduct(w http.ResponseWriter, r *http.Request)  {
    var product Product

   err := json.NewDecoder(r.Body).Decode(&product)

   if err != nil {
      writeJSON(w, http.StatusBadRequest, "Invalid Json")

      return
   }

   if product.Name == "" {
    writeJSON(w, http.StatusBadRequest, "Name is required")

    return
   }

   if product.Price <= 0 {
    writeJSON(w, http.StatusBadRequest, "Price must be greater than 0")

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
            writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "Unauthorized", Status: 401})
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

func writeJSON(w http.ResponseWriter, status int, data any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    return json.NewEncoder(w).Encode(data)
}


func getProduct(id int) (string, error){
  if id == 0 {
    err := fmt.Errorf("get product failed: %w", ErrProductNotFound)
    return "", err
  }

  return "Car", nil
}

func main() {

    mux := http.NewServeMux()

    mux.HandleFunc("GET /products", GetProduct)
    mux.HandleFunc("GET /products/{id}", GetProductDetail)

    mux.Handle("POST /products", AuthMiddleware(http.HandlerFunc(CreateProduct)))
    mux.Handle("PUT /products/{id}", AuthMiddleware(http.HandlerFunc(UpdateProduct)))
    mux.Handle("DELETE /products/{id}", AuthMiddleware(http.HandlerFunc(DeleteProduct)))

    handler := CORSMiddleware(LoggingMiddleware(mux))

   result, err := getProduct(0)

   if errors.Is(err, ErrProductNotFound){
     fmt.Println("produk tidak ditemukan")
   }else{
     fmt.Println(result)
   }

    http.ListenAndServe(":8080", handler)
}