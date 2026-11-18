package main

import (
	"log"
	"net/http"

	"go-project/internal/database"
	"go-project/internal/domain"
	"go-project/internal/handler"
	"go-project/internal/repository"
	"go-project/internal/usecase"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		log.Fatal(err)
	}

	// Repository
	productRepository := repository.NewProductRepository(db)

	// Use Case
	productUsecase := usecase.NewProductUsecase(productRepository)

	// Handler
	productHandler := handler.NewProductHandler(productUsecase)

	// Router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", productHandler.GetProducts)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProductByID)
	mux.HandleFunc("POST /products", productHandler.CreateProduct)

	log.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}