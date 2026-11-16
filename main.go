package main

import (
	"fmt"
)

func processTask(){

  defer fmt.Println("Cleanup 1")
  defer fmt.Println("Cleanup 2")

    defer func() {
        result := recover()

        if result != nil {
            fmt.Println("Recovered:", result)
        }
    }()

    fmt.Println("Task started")

    panic("task failed")
}

func main() {
  processTask()
  fmt.Println("Main continues")
}