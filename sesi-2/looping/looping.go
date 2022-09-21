package main

import "fmt"

func main() {

	//looping in GO
	//for loop first way
	for i := 0; i < 5; i++ {
		fmt.Println("Angka -", i, "\n")
	}

	//for loop second way
	j := 0
	for j < 5 {
		fmt.Println("Angkaa -", j, "\n")
		j++
	}

	//for loop third way
	var k = 0

	for {
		fmt.Println("Angkaaa -", k, "\n")

		k++
		if k == 5 {
			break
		}
	}

	fmt.Print("\n------------------------\n")
	//lopping with continue and break
	// saat bilangan ganjil diskip dengan "continue" dan saat bilangan lebih besar dari 8 diberhentikan dengan "break"
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	fmt.Print("\n------------------------\n")
	//nested loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	fmt.Print("\n------------------------\n")
	//labeling nested loop
outerLoop:
	for i := 0; i < 5; i++ {
		fmt.Print("\nperulangan i ke- [", i, "]", "\n")
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print(" perulangan j ke- [", j, "]")
		}
	}

}
