package main

import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Printf("int size of %+v\n", unsafe.Sizeof(int(9)))  //如果对size有明确要求，那就用 int32 之类的吧
    fmt.Printf("uint size of %+v\n", unsafe.Sizeof(uint(9)))
    fmt.Printf("int32 size of %+v\n", unsafe.Sizeof(int32(9)))
    fmt.Printf("int64 size of %+v\n", unsafe.Sizeof(int64(9)))

    sli:=make([]int,0)
    fmt.Printf("slice size of %+v\n", unsafe.Sizeof(sli))

    a := []int{}
    a = append(a, 7,8,9)
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

