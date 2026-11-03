package main

import "fmt"

type Product struct {
    Name string
    Price float64
    Stock int
    IsAvailable bool
}

func main() {

    product := Product{
        Name: "mACBOOK",
        Price: 12.500000,
        Stock: 12,
        IsAvailable: true,
    }

    product.Stock = 8

    fmt.Println(product)
}