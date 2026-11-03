package main

import "fmt"

func main() {

	age := new(int)

	*age = 100

	fmt.Println(age)
	fmt.Println(*age)
}
