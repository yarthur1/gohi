package main

import (
    "fmt"
    "unsafe"
)

func main() {
    a := []int{7,8,9} //a的cap=3,ap函数添加元素时会扩容申请一块新的内存，和slice a没关系
    fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
    ap(a)
    fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)
    p := unsafe.Pointer(&a[2])
    q := uintptr(p)+8
    t := (*int)(unsafe.Pointer(q))
    fmt.Println(*t)
}

func ap(a []int) {
    a = append(a, 10)
}


