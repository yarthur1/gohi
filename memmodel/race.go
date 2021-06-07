package main

import (
    "fmt"
    "time"
)
// https://golang.org/ref/mem  多线程一定要有同步事件
func main() {
    var x1 int
    go func() {
        for {
            time.Sleep(1 *time.Millisecond)
            x1++
        }
    }()
    time.Sleep(1 *time.Second)
    fmt.Println(x1)
}
