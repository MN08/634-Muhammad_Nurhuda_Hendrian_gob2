package main

import (
	"fmt"
	"os"
)

func main() {
	//defer
	fmt.Print("\n------------------------\n")
	// defer fmt.Println("defer execution started")
	fmt.Println("hai")
	fmt.Println("Welcome")

	//defer#2
	fmt.Print("\n\n------------------------\n")
	callDeferFunction()
	fmt.Println("hello world")

	// Exit
	fmt.Print("\n\n------------------------\n")
	defer fmt.Println("Invoke with defer")
	fmt.Println("before Exiting...")
	os.Exit(1)

}
func callDeferFunction() {
	fmt.Println("defer func start to execute")
}
