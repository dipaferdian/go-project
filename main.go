package main

import "fmt"

type Product struct {
    Name  string
    Price float64
    Stock int
}


func (p *Product) UpdateStock(stock int) int  {
    p.Stock = stock
    return p.Stock
}

func (p Product) GetInfo() string {
    p.Name = "Iphone"

    return p.Name
}

func main() {
    
  product := Product{
    Name: "Samsung",
    Price: 20,
    Stock: 2,
  }

    fmt.Println(product.GetInfo())
    fmt.Println(product.UpdateStock(1))
    fmt.Println(product.Stock)
}