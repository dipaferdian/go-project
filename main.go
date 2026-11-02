package main

import "fmt"

func main() {

	operation := func(name string) {
		fmt.Println("Hello ", name)
	}

	operation("dipa")
}
