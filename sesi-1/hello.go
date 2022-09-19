package main

import "fmt"

func main() {
	fmt.Println("hello World")
	fmt.Println("ini sesi pertama belajar golang \n")

	var fullName string = "mnurhudah"
	var age int = 24

	fmt.Println("Hello my name is ", fullName)
	fmt.Println("i am ", age, " yo \n")

	fullName = "dsda"
	fmt.Println("Hello my name is", fullName, "\n")
	fmt.Printf("%T", fullName, "\n")

	fmt.Println("=== variable tanpa var dengan value | short declaration ==== \n")
	firstName := "First Name"
	fmt.Println("This is my first name : ", firstName, "\n")

	fmt.Println("=== mutliple variable declaration ====\n")
	var nameOne, nameTwo, nameThree string = "name1", "name2", "name3"
	fmt.Println("This is multiple variable : ", nameOne, nameTwo, nameThree, "\n")

	//jika ada variable blm terpanggil
	fmt.Println("=== underscore variable ====\n")
	var someOne int
	var namesOne, namesTwo, namesThree string = "name1", "name2", "name3"
	_, _, _, _ = someOne, namesOne, namesTwo, namesThree
	fmt.Println("test underscore variable", namesOne, namesTwo, namesThree, "\n")

	//mengetahui penggunaan tipedata variable
	fmt.Println("=== fmt.printf usage ====\n")
	first, second := 1, "2"
	fmt.Printf("tipe data varibale first = %T", first, "\n")
	fmt.Printf("tipe data varibale second = %T", second)

	fmt.Println("=== Constanta ====\n")
	const constanta_name string = "this is contanta value\n"
	fmt.Println("print const", constanta_name)

	fmt.Println("=== Operator ====\n")
	fmt.Println("operator aritmatika")
	var num int
	num = 1 - 1 + 100
	fmt.Println(num, "\n")

	fmt.Println("operator Relational ")
	var moreThan bool = 10 > 1
	fmt.Println(moreThan)
}
