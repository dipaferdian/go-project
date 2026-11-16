package main

import (
	"encoding/json"
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

func main() {
    http.HandleFunc("GET /products", GetProduct)
    http.HandleFunc("POST /products", CreateProduct)

    http.ListenAndServe(":8080", nil)
}