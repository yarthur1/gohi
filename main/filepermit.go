package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.FileMode(777), 777)
    fmt.Println(os.FileMode(0777), 0777)
    fmt.Println(os.FileMode(0666))
    fmt.Println(0777)
}