package main

import "fmt"

func main() {
	//basic print command
	fmt.Println("hello World")
	fmt.Print("ini sesi pertama belajar golang \n")

	//variable with data type
	var fullName string = "mnurhudah"
	var age int = 24
	fmt.Print("\n------------------------\n")
	fmt.Println("Hello my name is ", fullName)
	fmt.Print("i am ", age, " yo \n")

	// variable without data type
	fmt.Print("\n------------------------\n")
	fullName = "And dhik akang enband"
	fmt.Println("Hello my name is ", fullName, "\n")
	//Printf usage
	fmt.Printf("%s", fullName, "\n")

	//variabel without datatype short Declaration
	firstName := "First Name"
	fmt.Println("This is my first name : ", firstName, "\n")

	//Multiple Variabel Declaration
	var nameOne, nameTwo, nameThree string = "name1", "name2", "name3"
	fmt.Println("This is multiple variable : ", nameOne, nameTwo, nameThree, "\n")

	//jika ada variable blm terpanggil
	//underscore variable
	var someOne int
	var namesOne, namesTwo, namesThree string = "name1", "name2", "name3"
	_, _, _, _ = someOne, namesOne, namesTwo, namesThree
	fmt.Println("\n------------------------\n")
	fmt.Println("test underscore variable", namesOne, namesTwo, namesThree, "\n")

	//mengetahui penggunaan tipedata variable
	//fmt.printf usage
	first, second := 1, "2"
	fmt.Println("\n------------------------\n")
	fmt.Printf("tipe data varibale first = %T \n", first)
	fmt.Printf("tipe data varibale second = %T \n", second)

	//Constanta
	const constanta_name string = "this is contanta value\n"
	fmt.Println("\n------------------------\n")
	fmt.Println("print const", constanta_name)

	// Operator
	fmt.Print("\n------------------------\n")
	fmt.Println("operator aritmatika")
	var num int
	num = 1 - 1 + 100
	fmt.Println(num, "\n")

	fmt.Print("\n------------------------\n")
	fmt.Println("operator Relational ")
	var moreThan bool = 10 > 1
	fmt.Println(moreThan)
}
