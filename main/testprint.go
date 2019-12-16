package main

import "fmt"

func main(){
	a := 10
	fmt.Println(a)  //right
	fmt.Println("abc") //right
	fmt.Printf("%d", a)//right
	fmt.Printf(a) //error
}
