package main
import (
    "fmt"
    "github.com/shopspring/decimal"
)
// https://blog.csdn.net/LinHenk/article/details/88821297

func main(){
    fmt.Println(int64(8.04758523075822e+17))  //科学计数法可以直接转化为int64

    var newNum float64
    numStr := "8.04758523075822e+17"          //科学计数法字符串
    _, err := fmt.Sscanf(numStr, "%e", &newNum)
    if err != nil {
        return
    }
    fmt.Println(newNum)

    fmt.Println(int64(newNum))      //精度变小
    num := fmt.Sprintf("%.f", newNum)
    fmt.Println(num)

    decimalNum, err := decimal.NewFromString(numStr)
    if err != nil {
        return
    }
    fmt.Println(decimalNum.String())   //精度不变

}
