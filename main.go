package main

import (
	"errors"
	"fmt"
)

func getProduct(id int) (string, error) {

	if id == 0 {
		return "", errors.New("product not found")
	} else {
		return "Car", nil
	}
}

func main() {

	ok, err := getProduct(1)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ok)
	}

}
