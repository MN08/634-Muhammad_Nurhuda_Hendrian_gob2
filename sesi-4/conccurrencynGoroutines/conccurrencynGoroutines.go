package main

import "fmt"

func main() {

	// Concurrency
	fmt.Print("\n------------------------\n")
	go goroutine()

	// Goroutines (Asynchronous process #1)
	fmt.Print("\n\n------------------------\n")

}

func goroutine() {
	fmt.Println("hello")
}
