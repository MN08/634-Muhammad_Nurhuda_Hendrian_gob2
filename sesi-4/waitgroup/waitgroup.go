package main

import (
	"fmt"
	"sync"
)

func main() {

	// wait group go routines
	fmt.Print("\n------------------------\n")
	fruits := []string{"apple", "jack fruit", "orange", "banana", "avocado"}
	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruits(index, fruit, &wg)
	}

	wg.Wait()

}

func printFruits(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index : %d, fruit : %s\n", index, fruit)
	wg.Done()
}
