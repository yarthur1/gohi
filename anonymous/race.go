package main

import (
    "fmt"
    "time"
)

func main() {
    var x1 int
    go func() {
        for {
            x1++
        }
    }()
    time.Sleep(1 *time.Second)
    fmt.Println(x1)
}
