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

//https://juejin.cn/post/6844903679519096846

func returnValues() int {  //没有命名返回值，修改不会保留
    var result int
    defer func() {
        result++
        fmt.Println("defer")
    }()
    return result
}

func namedReturnValues() (result int) {
    defer func() {
        result++
        fmt.Println("defer")
    }()
    return result
}


func main(){
    doSomething()
    fmt.Println(doSomething1())
    fmt.Println(doSomething2())
    fmt.Println(returnValues())
    fmt.Println(namedReturnValues())
}
