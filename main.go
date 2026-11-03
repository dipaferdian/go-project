package main

import (
	"errors"
	"fmt"
)

func getProduct(id int) (string, error) {
    if id == 0 {
        return "", errors.New("product not found")
    }

    return "", nil
}

func getProductService(id int) (string, error) {
    product, err := getProduct(id)

    if err != nil {
        return "", fmt.Errorf("get product service: %w", err)
    }

    return product, nil
}

func main() {

	ok, err := getProductService(0)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ok)
	}

}
