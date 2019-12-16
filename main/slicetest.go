package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice_1 := make([]int, 5, 6)
	// slice_1 = []int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		slice_1[i] = i + 1
	}
	sliceHeader := unsafe.Pointer(&slice_1[4])
	fmt.Printf("slice_1-->data:\t%#v\n", slice_1)
	//fmt.Printf("main-->ptr:\t%#v\n", sliceHeader.Data)
	fmt.Printf("slice_1-->len:\t%#v\n", len(slice_1))
	fmt.Printf("slice_1-->cap:\t%#v\n", cap(slice_1))
	test1(slice_1)
	fmt.Printf("slice_1-->data:\t%#v\n", slice_1)
	fmt.Printf("slice_1-->len:\t%#v\n", len(slice_1))
	fmt.Printf("slice_1-->cap:\t%#v\n", cap(slice_1))

	fmt.Printf("yxj-->ptr:\t%#v\n", uintptr(sliceHeader)+8)
	t := (*int)(unsafe.Pointer(uintptr(sliceHeader)+8))
	fmt.Printf("yxj-->data:\t%#v\n",*t)
	test2(&slice_1)
	fmt.Printf("main-->data:\t%#v\n", slice_1)

}

func test1(slice_2 []int) {
	slice_2[1] = 6666               //函数外的slice确实有被修改
	slice_2 = append(slice_2, 8888) //函数外的不变
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice_2))
	fmt.Printf("slice_2-->ptr:\t%#v\n", sliceHeader.Data)
	fmt.Printf("slice_2-->data:\t%#v\n", slice_2)
	fmt.Printf("slice_2-->len:\t%#v\n", len(slice_2))
	fmt.Printf("slice_2-->cap:\t%#v\n", cap(slice_2))
}

func test2(slice_2 *[]int) { //这样才能修改函数外的slice
	*slice_2 = append(*slice_2, 6666)
}