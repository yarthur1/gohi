package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    var count int
    var ch = make(chan bool, 1)
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            ch <- true
            count++
            fmt.Println(count)
            time.Sleep(time.Second)
            count--
            fmt.Println(count)
            <-ch
            wg.Done()
        }()
    }
    wg.Wait()
}