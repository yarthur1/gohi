package main

import "fmt"

func functions() []func() {
	// pitfall of using loop variables
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	for i := range arr {
		result = append(result, func() { fmt.Printf("index - %d, value - %d\n", i, arr[i]) })
	}

	return result
}

func main() {
	fns := functions()
	for f := range fns {
		fns[f]()
	}
}
