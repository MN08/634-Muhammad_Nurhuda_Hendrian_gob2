package main

import (
	"fmt"
	"os"
)

func main() {
	//defer
	fmt.Print("\n------------defer------------\n")
	defer fmt.Println("defer execution started")
	fmt.Println("hai")
	fmt.Println("Welcome")

	//defer#2
	fmt.Print("\n\n------------defer#2------------\n")
	callDeferFunction()
	fmt.Println("hello world")

	// Exit
	fmt.Print("\n\n------------Exit------------\n")
	defer fmt.Println("Invoke with defer")
	fmt.Println("before Exiting...")
	os.Exit(1)

}
func callDeferFunction() {
	fmt.Println("defer func start to execute")
}
