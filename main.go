package main

import "fmt"
type User struct {
    Name string
    Age int
}

func main() {
    var user *User

    if user == nil {

        user = &User{
            Name: "Mac",
            Age: 12,
        }

        fmt.Println(*user)
    }
}