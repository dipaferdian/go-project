package main

import (
	"errors"
	"fmt"
)

var ErrProductNotFound = errors.New("product not found")

func getProduct(id int) (string, error) {
	if id == 0 {
		return "", ErrProductNotFound
	}

	return "Car", nil
}

func getProductService(id int) (string, error) {
	product, err := getProduct(id)

	if err != nil {
		return "", fmt.Errorf("get product service: %w", err)
	}

	return product, nil
}

func main() {

	product, err := getProductService(0)

	if errors.Is(err, ErrProductNotFound) {
		fmt.Println("Product tidak ditemukan")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(product)
	}

}
