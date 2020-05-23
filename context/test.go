package main

import (
    "fmt"
    "time"
    "context"
)

func main() {
    ctx := context.Background()
    ch := make(chan struct{})
    go func(){
       time.Sleep(1*time.Second)
        ch <-struct{}{}
    }()
    select {
    case <-ctx.Done():
       fmt.Printf("background done")
       return
    case <-ch:
      fmt.Println("ch done")
    default:
        fmt.Println("default done")
    }
    if ctx.Done() == nil {
      fmt.Println("nil")
    }
}
