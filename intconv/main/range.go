package main

import (
    "fmt"
    "strconv"
)

func isIntegerType(val []byte) (bool,int64){    //is integer ,return val,or return strlen
    l:=len(val)
    if l==0{
        return false,0   //index out of range [0] with length 0   call sizeOfString
    }
    if val[0]=='0' && l>1{
        return false,int64(l)
    }
    num,err:=strconv.ParseInt(string(val),10,64)   //REDIS_ENCODING_INT long
    if err==nil{
        return true,num
    }
    return false,int64(l)
}

func main(){
    st:="9223372036854775808"
    var sl=[]byte(st)
    flag,_:=isIntegerType(sl)
    if flag{
        fmt.Println("true")
    }else{
        fmt.Println("false")
    }
}
