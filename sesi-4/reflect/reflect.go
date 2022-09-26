package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Identifying Data Type & Value
	fmt.Print("\n------------------------\n")
	var number = 20
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("type variable number : ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {

		fmt.Println("number value : ", reflectValue.Int())
	}

	//Accessing Value Using Interface Method
	fmt.Print("\n\n------------------------\n")
	var numb = 40
	var reflectValues = reflect.ValueOf(numb)

	var nilai = reflectValues.Interface().(int)
	fmt.Println("tipe  variabel :", reflectValues.Type())
	fmt.Println("nilai variabel :", reflectValues.Interface())
	fmt.Println("nilai variabel :", nilai)

	//Identifying Method Information
	fmt.Print("\n\n------------------------\n")
	var s1 = &student{Name: "Student satu", grade: 4}
	fmt.Println("name :", s1.Name)

	var reflectValuee = reflect.ValueOf(s1)
	var method = reflectValuee.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("satu"),
	})

	fmt.Println("name :", s1.Name)

}

type student struct {
	Name  string
	grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}
