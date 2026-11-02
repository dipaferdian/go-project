package main

import "fmt"

func main() {

    greet := func(name string) string {
        return "Hello " + name
    }

    result := execute(greet)

	fmt.Println(result)
}


func execute(callback func(string) string) string {
    return callback("dipa")
}