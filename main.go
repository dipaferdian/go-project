package main

import "fmt"

type Product struct {
    Name string
    Price float64
    Stock int
    IsAvailable bool
}

func printProduct(product Product)  {
    fmt.Println(product)
}

func main() {

    product := Product{
        Name: "mACBOOK",
        Price: 12.500000,
        Stock: 12,
        IsAvailable: true,
    }

    printProduct(product)
}