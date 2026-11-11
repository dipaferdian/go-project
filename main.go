package main

import (
	"fmt"
	"sync"
)


func sendData(ch chan<- int, wg *sync.WaitGroup){
  defer wg.Done()

  ch <- 20
  ch <- 30
  ch <- 40

  close(ch)
}

func receiveData(ch <-chan int, wg *sync.WaitGroup)  {
  defer wg.Done()

  for value := range ch {
    fmt.Println(value)
  }
}

func main() {
  var wg sync.WaitGroup

  wg.Add(2)

  ch := make(chan int, 3)

  go sendData(ch, &wg)
  go receiveData(ch, &wg)

  wg.Wait()
}