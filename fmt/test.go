package main

import (
    "fmt"

)

func main() {
    sqlGetFuids := "SELECT fuid FROM %s WHERE uid=? AND fuid in (%s)"
    tmp := fmt.Sprintf(sqlGetFuids, "table")  //差一个参数
    fmt.Println(tmp)
    fmt.Println(fmt.Sprintf(tmp, "table"))

}
