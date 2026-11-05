package main

import (
	"fmt"
	"go-project/user"
)
func main() {
    
    result := user.NewUser("Dipa", 29)

    fmt.Println(result)
}