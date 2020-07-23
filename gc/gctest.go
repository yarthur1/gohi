package main

import (
    "time"
)

func main() {

    for {
        str := "qwwe"
        for i := 0; i < 10000; i++ {
            str += "sadhjjfjsjhjdjjs"
        }

        chanWaitGroup := make(chan string, 1)

        // ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

        go func() {
            time.Sleep(200 * time.Millisecond)
            chanWaitGroup <- str
        }()

        select {
        case <-chanWaitGroup:
            close(chanWaitGroup)
            //case <-ctx.Done():
            //    cancel()
        default:

        }
    }
}
