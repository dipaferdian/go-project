package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup){
  defer wg.Done()

  select {
  case <-ctx.Done():
    fmt.Println("Worker stopped")
  case <-time.After(3 * time.Second):
    fmt.Println("Worker completed")
  }
}

func main() {
  var wg sync.WaitGroup

  wg.Add(1)

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  go worker(ctx, &wg)

  time.Sleep(1 * time.Second)

  cancel()

  wg.Wait()
}