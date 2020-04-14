package main

import "fmt"

func doSomething() {
    v := 10
    defer func() {
        fmt.Println(v)
        v++
        fmt.Println(v)
    }()
    v += 5
}

func doSomething1() (rev int) {
    defer func() {  //修改外围函数返回值
        rev++
        fmt.Println(rev)
    }()

    return 5  //rev = 5  defer return rev
}

func doSomething2() (rev int) {
    v := 10
    defer func() {
        v++
        fmt.Println(rev)
        rev++
        fmt.Println(v)
    }()

    return v  //rev = v  defer return rev
}

func main(){
    doSomething()
    fmt.Println(doSomething1())
    fmt.Println(doSomething2())
}
