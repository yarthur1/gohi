package main

import "fmt"

func main() {
    // 定义map
    map1 := map[int]string{1: "张三", 2: "李四", 3: "王五"}
    fmt.Println(map1)  // map[1:张三 2:李四 3:王五]
    fmt.Printf("%p \n", map1)  // 0xc000060240

    // 调用函数  map作为函数参数
    test(map1)
    fmt.Println(len(map1))
}

// map作为函数的参数
func test(m map[int]string) {
    fmt.Printf("%p \n", m)  // 0xc000060240 形参的地址与实参相同（引用传递）
    // 修改形参的值，会影响到实参
    m[1] = "哈哈哈"  // 修改
    m[10] = "啦啦啦"  // 添加
    for i:=100;i<10000;i++{
        m[i]="test"
    }
    fmt.Printf("after mod %p \n", m)
}
