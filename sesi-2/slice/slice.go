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

	//slice (slicing)

	fmt.Print("\n------------------------\n")
	var fruuits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruuits[0:3]
	var bFruits = fruuits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruuits)  // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruuits)  // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	//Slice (Combining slicing and append)

	fmt.Print("\n------------------------\n")
	var fruitsss1 = []string{"apple", "pineapple", "banana"}
	fruitsss1 = append(fruitsss1[:2], "papaya")
	fmt.Printf("%#v", fruitsss1)

	// Slice (Backing array)
	fmt.Print("\n------------------------\n")
	var fruitssss = []string{"apple", "pineapple", "banana", "papaya"}
	var fruitssss1 = fruitssss[1:3]

	fruitssss1[1] = "grapes"

	fmt.Println("fruitsssss = ", fruitssss)
	fmt.Println("fruitsssss1 = ", fruitssss1)

	//Slice (Cap function)
	fmt.Print("\n------------------------\n")
	var ffruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(cap(ffruits)) // cap: 4
	fmt.Println(len(ffruits)) // len: 4

	var aFfruits = ffruits[0:3]
	fmt.Println(cap(aFfruits)) // cap: 4
	fmt.Println(len(aFfruits)) // len: 3

	var bFfruits = ffruits[1:4]
	fmt.Println(cap(bFfruits)) // cap: 3
	fmt.Println(len(bFfruits)) // len: 3

	// Slice - Sesi 2 Slice (Creating a new backing array)
	fmt.Print("\n------------------------\n")
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	cars[0] = "nissan"
	fmt.Println("cars : ", cars)
	fmt.Println("newCars : ", newCars)

}
