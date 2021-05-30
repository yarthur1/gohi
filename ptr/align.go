package main

import (
    "fmt"
    "unsafe"
)

type SizeOfC struct {
    A byte  // 1字节
    C int32 // 4字节
}

type SizeOfD struct {
    A byte
    B [3]int32
}

type SizeOfE struct {
    A byte  // 1
    B int64 // 8
    C byte  // 1
}

type SizeOfF struct {
    A byte  // 1
    C byte  // 1
    B int64 // 8
}
//https://www.cnblogs.com/-wenli/p/12681044.html
func main() {
    fmt.Printf("%+v\n",unsafe.Sizeof(SizeOfC{0, 0}))
    fmt.Printf("%+v\n",unsafe.Alignof(SizeOfC{0, 0}))

    fmt.Printf("%+v\n",unsafe.Sizeof(SizeOfD{}))
    fmt.Printf("%+v\n",unsafe.Alignof(SizeOfD{}))

    fmt.Printf("%+v\n",unsafe.Sizeof(SizeOfE{}))
    fmt.Printf("%+v\n",unsafe.Alignof(SizeOfE{}))

    fmt.Printf("%+v\n",unsafe.Sizeof(SizeOfF{}))
    fmt.Printf("%+v\n",unsafe.Alignof(SizeOfF{}))

}



