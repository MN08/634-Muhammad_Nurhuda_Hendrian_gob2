package main

import (
	"fmt"
	"strings"
)

func main() {
	//POINTER
	//Pointer (Memory address)
	fmt.Print("\n------------------------\n")
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber value= ", firstNumber)
	fmt.Println("firstNumber memory= ", &firstNumber)

	fmt.Println("secondNumber value= ", *secondNumber)
	fmt.Println("secondNumber memori= ", secondNumber)

	//Pointer (Changing value through pointer)
	fmt.Print("\n------------------------\n")
	var firstPerson string = "jhon"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson value= ", firstPerson)
	fmt.Println("firstPerson memory= ", &firstPerson)
	fmt.Println("secondPerson value= ", *secondPerson)
	fmt.Println("secondPerson memory= ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson value= ", firstPerson)
	fmt.Println("firstPerson memory= ", &firstPerson)
	fmt.Println("secondPerson value= ", *secondPerson)
	fmt.Println("secondPerson memory= ", secondPerson)

	//Pointer (Pointer as a parameter)
	fmt.Print("\n------------------------\n")
	var a int = 10
	fmt.Println("before :", a)
	changeValue(&a)

	fmt.Println("after value=", a)

}

func changeValue(number *int) {
	*number = 300
}
