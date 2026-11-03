package main

import "fmt"

type Product struct {
    Name string
    Stock int
}

func updateStock(product *Product)  {
    product.Stock = 20
}

func main() {

    product := Product{
        Name: "Mac",
        Stock: 12,
    }

    updateStock(&product)

    fmt.Println(product.Stock)
}