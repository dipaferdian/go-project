package main

import "fmt"

func main() {
	role := "admin"

	switch role {
	case "admin":
		fmt.Println("Full Access")
	case "editor":
		fmt.Println("Can Edit")
	case "viewer":
		fmt.Println("Read Only")
	default:
		fmt.Println("Unknown role")
	}
}
