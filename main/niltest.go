package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%p\n", sl)
	fmt.Printf("%p\n", f)
	fmt.Printf("%#v\n", i)
}