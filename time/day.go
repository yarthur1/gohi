package main

import (
    "fmt"
    "time"
)

func main() {
    d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
    day := d.Day()

    fmt.Printf("day = %v\n", day)

    fmt.Printf(time.Unix(0, -100000).Format("2006-01-02 15:04:05"))

}
