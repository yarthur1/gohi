package main

import (
    "sync"
    "time"
)

var wg sync.WaitGroup

func work() {
    for {
        wg.Add(1)
        // ... work here
        time.Sleep(time.Second)
        wg.Done()
    }
}

func main() {
    go work()
    defer wg.Wait()
}