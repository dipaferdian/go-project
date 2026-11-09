package main

import (
	"fmt"
	"sync"
)


type Task struct {
  Name string
}

func (task Task) Run(wg *sync.WaitGroup) {
  defer wg.Done()

  fmt.Println("Running:", task.Name)
}

func main() {
    var wg sync.WaitGroup

    wg.Add(3)

    task1 := Task{
      Name: "Send Email",
    }

    task2 := Task{
      Name: "Process Payment",
    }

    task3 := Task{
      Name: "Generate Report",
    }

    go task1.Run(&wg)
    go task2.Run(&wg)
    go task3.Run(&wg)


    wg.Wait()

    fmt.Println("All tasks finished")
}