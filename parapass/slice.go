package main

import (
    "fmt"
    "unsafe"
)

func main() {
    a := []int{7,8,9}
    fmt.Printf("%p \n", a)  //slice   %p address of 0th element in base 16 notation, with leading 0x
    fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
    ap(a)
    fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
    p := unsafe.Pointer(&a[2])
    q := uintptr(p)+8
    t := (*int)(unsafe.Pointer(q))
    fmt.Println(*t)
}

func ap(a []int) {
    fmt.Printf("%p \n", a) //内存地址没变
    a = append(a, 10)
    fmt.Printf("%p \n", a) //内存地址已改变
    fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
}

