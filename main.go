package main

import (
	"fmt"
	"time"
)


func printMessage(message string)  {
    fmt.Println(message)
}

func main() {
  go printMessage("test1")
  go printMessage("test2")
  go printMessage("test2")

  time.Sleep(1 * time.Second)

  fmt.Println("Main finished")
}