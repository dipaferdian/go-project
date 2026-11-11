package main

import "fmt"

func sendData() chan int {
    ch := make(chan int, 4)

    ch <- 10
    ch <- 20
    ch <- 30
    ch <- 40

    close(ch)

    return ch
}

func main() {

  result := sendData()

  for value := range result {
    fmt.Println(value)
  }
}