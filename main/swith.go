package main

import "fmt"

func main() {

    switch false{
    case false:
        fmt.Println("The integer was <= 4")
        fallthrough
    case true:
        fmt.Println("The integer was <= 5")
        fallthrough
    case false:
        fmt.Println("The integer was <= 6")
        fallthrough
    case true:
        fmt.Println("The integer was <= 7")
        fallthrough
    case false:
        fmt.Println("The integer was <= 8")
        fallthrough
    default:
        fmt.Println("default case")
    }
}