package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now().Format("2006-01-02"))
	var DeliverAt time.Time
	fmt.Println(DeliverAt.UnixNano())
}