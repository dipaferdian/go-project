package main

import "fmt"

func main() {

	var age *int

    if age == nil {
        fmt.Println("Pointer is nil")
    }else{
        fmt.Println(*age)
    }
}
