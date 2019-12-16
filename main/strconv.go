package main

import (
    "fmt"
    "strconv"
)

func main(){
    str:="0123333333333333333333333333333333333333333333333333333333333333"
    num,err:=strconv.ParseInt(str,10,64)   //REDIS_ENCODING_INT long
    if err==nil{
        fmt.Println(num)
    }
}