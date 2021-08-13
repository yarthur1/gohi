package main

import (
    "encoding/hex"
    "fmt"
    "unicode/utf8"
)

func main() {

    var str = "ac"
    fmt.Println("len count:" , len(str))
    hex_string:=hex.EncodeToString([]byte(str))
    fmt.Println("hex string:" , hex_string)
    fmt.Println("hex len count:" , len(hex_string))

    str = "中国"
    hex_string=hex.EncodeToString([]byte(str))
    fmt.Println("len count:" , len(str))
    fmt.Println("hex string:" , hex_string)
    fmt.Println("hex len count:" , len(hex_string))
    fmt.Println("***********************")

    var str1 =  "̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̨̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍̍"   //并不是一个特殊字符，由多个特殊字符组成,单个字符最多4个字节
    fmt.Printf("%v\n",str1)
    fmt.Println("rune count:" , utf8.RuneCountInString(str1))
    fmt.Println("len count:" , len(str1))
    hex_string_data := hex.EncodeToString([]byte(str1))  // 转成16进制字符串
    fmt.Println("hex len count:" , len(hex_string_data))
    fmt.Printf("%v\n",hex_string_data)

    var str2 =  "2̨̨̨̨̨̨̨̨̨̨̍̍̍̍̍̍̍̍̍"
    fmt.Printf("%v\n",str2)
    fmt.Println("rune count:" , utf8.RuneCountInString(str2))
    fmt.Println("len count:" , len(str2))
    hex_string_data2 := hex.EncodeToString([]byte(str2))
    fmt.Println("hex len count:" , len(hex_string_data2))
    fmt.Println(hex_string_data2)

    var str3 =  "2̨̨̨̨̨̨̨̨̨̨̍̍̍̍̍̍̍̍̍ 2̨̨̨̨̨̨̨̨̨̨̍̍̍̍̍̍̍̍̍ 2̨̨̨̨̨̨̨̨̨̨̍̍̍̍̍̍̍̍̍"
    fmt.Printf("%v\n",str3)
    fmt.Println("rune count:" , utf8.RuneCountInString(str3))
    fmt.Println("len count:" , len(str3))
    hex_string_data3 := hex.EncodeToString([]byte(str3))
    fmt.Println("hex len count:" , len(hex_string_data3))
    fmt.Println(hex_string_data3)
}