package main
import (
    "fmt"
    "github.com/shopspring/decimal"
    "strconv"
)
// https://blog.csdn.net/LinHenk/article/details/88821297

func main(){
    fmt.Println(int64(8.04758523075822e+17))  //科学计数可以直接转化为int64  科学计数法字符串不能直接转int64
    numStr := "8.04758523075822e+17"          //科学计数法字符串

    newNum1, err := strconv.ParseInt(numStr, 10, 64)   //科学计数法字符串不能直接转int64
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(newNum1)

    var newNum float64

    _, err = fmt.Sscanf(numStr, "%e", &newNum)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(newNum)

    fmt.Println(int64(newNum))      //精度变小
    num := fmt.Sprintf("%.f", newNum)
    fmt.Println(num)

    decimalNum, err := decimal.NewFromString(numStr)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(decimalNum.String())   //精度不变

    newNum2, err := strconv.ParseFloat(numStr, 64)   //科学计数法字符串可以直接转float64,精度会变小
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(newNum2)
    fmt.Println(int64(newNum2))

}
