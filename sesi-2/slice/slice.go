package main

import "fmt"

func main() {

	//slice
	fmt.Print("\n------------------------\n")
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	//slice make function
	var fruit = make([]string, 3)
	_ = fruit
	fmt.Printf("\n%#v", fruit)

	//slice append function
	fmt.Print("\n------------------------\n")
	var fruitt = make([]string, 3)
	_ = fruitt
	fruitt[0] = "banana"
	fruitt[1] = "apple"
	fruitt[2] = "pineapple"
	fmt.Printf("\n%#v", fruitt)

	var fruiit = make([]string, 3)
	fruiit = append(fruiit, "banana", "apple", "pineapple")
	fmt.Printf("\n%#v", fruiit)

	//Slice (append function with ellipsis)
	fmt.Print("\n------------------------\n")
	var fruits1 = []string{"apple", "Pinapple", "MangoRose"}
	var fruits2 = []string{"jack fruit", "banana", "Orange"}

	fruits1 = append(fruits1, fruits2...)
	fmt.Println("\n", fruits1)

	// Slice (copy function)
	fmt.Print("\n------------------------\n")
	var fruitss1 = []string{"apple", "Pinapple", "MangoRose"}
	var fruitss2 = []string{"jack fruit", "banana", "Orange"}

	fmt.Println(fruitss1)
	fmt.Println(fruitss2)
	nn := copy(fruitss1, fruitss2)
	fmt.Println(fruitss1)
	fmt.Println(fruitss2)
	fmt.Println(nn)
}
