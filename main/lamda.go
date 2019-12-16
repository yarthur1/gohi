package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func getPrintMessage() func(string) {
	// returns an anonymous function
	return func(message string) {
		fmt.Println(message)
	}
}

func main() {
	// named function
	printMessage("Hello function!")

	// anonymous function declared and called
	func(message string) {
		fmt.Println(message)
	}("Hello anonymous function!")

	// gets anonymous function and calls it
	printfunc := getPrintMessage()
	printfunc("Hello anonymous function using caller!")

}
