package main

import (
	"fmt"
	"sync"
)

type Inventory struct {
  Stock int
  mu sync.Mutex
}

func (inventory *Inventory) AddStock(amount int, wg *sync.WaitGroup){
  inventory.mu.Lock()
  defer wg.Done()
  defer inventory.mu.Unlock()

  inventory.Stock += amount
}

func main() {
    var wg sync.WaitGroup  

    wg.Add(5)
    inventory := &Inventory{}

    go inventory.AddStock(10, &wg)
    go inventory.AddStock(10, &wg)
    go inventory.AddStock(10, &wg)
    go inventory.AddStock(10, &wg)
    go inventory.AddStock(10, &wg)

    wg.Wait()

    fmt.Println(inventory.Stock)
}