package main

import (
	"fmt"
	"strings"
)

func judgeTrie(str string) bool{
	if str==""{
		return false
	}
	for i:=0;i<len(str);i++{
		if str[i]=='\\' || str[i]=='?' || str[i]=='{' || str[i]=='}' ||
			str[i]=='(' || str[i]==')' || str[i]=='[' || str[i]==']' ||
			str[i]=='.' || str[i]=='*' || str[i]=='+' || str[i]=='|' {
			return false
		}
	}
	return true
}

func strFormat(str string) string{
	str=strings.Replace(str,"\\.",".",-1)
	l:=len(str)
	if l==0{
		return ""
	}
	if str==".*" || str[0]!='^' {
		return ""
	} else if str[l-1]=='$'{
		if judgeTrie(str[1:l-1]){
			return string([]byte(str[1:l-1]))
		}
		return ""
	} else if str[l-1]=='*' && str[l-2]=='.'{
		if judgeTrie(str[1:l-2]){
			return string([]byte(str[1:l-2]))
		}
		return ""
	}else if judgeTrie(str[1:]){
		return string([]byte(str[1:]))
	}
	return ""
}

func main(){
	first := "fisrt"
	str:=[]byte(first)
	if str[0]==102{
		fmt.Println(str)
	}

	//fmt.Println([]rune(first))
	//fmt.Println([]byte(first))
	str1:="^cos\\.graph\\.friend\\.\\d*$"
	if strFormat(str1)==""{
		fmt.Println("pass")
	}
	fmt.Println(str1)
	//全部替换
	fmt.Println(strings.Replace("", "oink", "moo", -1))
}
