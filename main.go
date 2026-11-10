package main

import "fmt"


func main() {

  ch := make(chan string, 2)

  ch <- "test1"
  ch <- "test2"

  fmt.Println(len(ch))
  fmt.Println(cap(ch))

  fmt.Println(<-ch)
  fmt.Println(<-ch)

  fmt.Println(len(ch))
  fmt.Println(cap(ch))


}