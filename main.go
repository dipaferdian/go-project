package main

import "fmt"

func main() {

	area, perimeter := calculateRectangle(10, 5)

	fmt.Println("area:", area)
	fmt.Println("perimeter:", perimeter)

}

func calculateRectangle(width int, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)

	return
}
