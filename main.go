package main

import "fmt"

func main() {
	users := map[int]string{
		1: "Dipa",
		2: "Andi",
		3: "Budi",
	}

	fmt.Println(users[3])

	delete(users, 2)

	users[4] = "Rudi"

	for _, value := range users {
		fmt.Println(value)
	}

}
