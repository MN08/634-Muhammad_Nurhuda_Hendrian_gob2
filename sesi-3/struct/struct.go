package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name string
	age  int
}

type Structure struct {
	division string
	person   Person
}

// syntax
type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	//STRUCT
	//syntax
	// Struct (Giving value to struct)
	fmt.Print("\n------------------------\n")
	var employee Employee
	employee.name = "John"
	employee.age = 43
	employee.division = "it"

	fmt.Println("name : ", employee.name)
	fmt.Println("age : ", employee.age)
	fmt.Println("division : ", employee.division)

	// Struct (Initializing struct)
	fmt.Print("\n------------------------\n")
	var employee1 = Employee{}
	employee1.name = "aJohns"
	employee1.age = 42
	employee1.division = "it"

	var employee2 = Employee{name: "John", age: 32, division: "it"}

	fmt.Printf("Employee1 : %#v\n ", employee1)
	fmt.Printf("Employee2 : %#v\n ", employee2)

	//Struct (Pointer to a struct)
	fmt.Print("\n------------------------\n")
	var employee3 = Employee{name: "John", age: 32, division: "it"}
	var employeePointer *Employee = &employee3

	fmt.Println(strings.Repeat("#", 25))

	fmt.Println("Employee3 name before : ", employee3.name)
	employeePointer.name = "alberto"

	fmt.Println("Employee3 name now: ", employee3.name)
	fmt.Println("EmployeePointer name : ", employeePointer.name)

	//Struct (Embedded struct)
	fmt.Print("\n------------------------\n")
	var structure = Structure{}
	structure.person.name = "agus"
	structure.person.age = 32
	structure.division = "it"

	fmt.Printf("%+v", structure)

	// /Struct (Anonymous struct)
	fmt.Print("\n\n------------------------\n")
	var anonymousStruct = struct {
		person   Person
		division string
	}{}
	anonymousStruct.person = Person{name: "alberto", age: 32}
	anonymousStruct.division = "it"

	var anonymousStruct2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "albertus", age: 43},
		division: "HR",
	}

	fmt.Printf("%+v", anonymousStruct)
	fmt.Println()
	fmt.Printf("%+v", anonymousStruct2)

	//Struct (Slice of struct)
	fmt.Print("\n\n------------------------\n")
	var people = []Person{
		{name: "albert", age: 34},
		{name: "alberto", age: 34},
		{name: "albertus", age: 34},
	}

	for _, v := range people {
		fmt.Printf("%+v", v)
		fmt.Println()
	}

	//Struct (Slice of anonymous struct)
	fmt.Print("\n\n------------------------\n")

	var karYawan = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "albertus", age: 34}, division: "it"},
		{person: Person{name: "alberto", age: 34}, division: "it"},
		{person: Person{name: "albert", age: 34}, division: "it"},
	}

	for _, v := range karYawan {
		fmt.Printf("%+v", v)
		fmt.Println()
	}
}
