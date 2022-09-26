package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// Concurrency
	fmt.Print("\n------------------------\n")
	// go goroutine()

	// Goroutines (Asynchronous process #1)
	fmt.Print("\n\n------------------------\n")
	fmt.Println("main execution start")

	go firstProcess(8)
	secondProcess(8)

	fmt.Println("no of Goroutines : ", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)

	fmt.Println("main execution stop")

}

// func goroutine() {
// 	fmt.Println("hello")
// }

// Goroutines (Asynchronous process #1)
func firstProcess(index int) {
	fmt.Println("First process func start")
	for i := 1; i < index; i++ {
		fmt.Println("i = ", i)
	}
	fmt.Println("First process func stop")
}
func secondProcess(index int) {
	fmt.Println("Second process func start")
	for j := 1; j < index; j++ {
		fmt.Println("j = ", j)
	}
	fmt.Println("Second process func stop")
}
