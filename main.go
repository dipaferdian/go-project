package main

import "fmt"


func sendNumbers(ch chan int, value int) {
    ch <- value
}

func main() {

  ch := make(chan int)

  go sendNumbers(ch, 10)
  go sendNumbers(ch, 20)
  go sendNumbers(ch, 30)


  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch)

}