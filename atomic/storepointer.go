package main

import (
    "fmt"
    "unsafe"
)
import "sync/atomic"

func main() {
    var a int =4
    var b int =7
    p :=unsafe.Pointer(&a)
    fmt.Printf("%#v\n",*(*int)(p))
    atomic.StorePointer(&(p), unsafe.Pointer(&b))
    fmt.Printf("%#v\n",*(*int)(p))
}