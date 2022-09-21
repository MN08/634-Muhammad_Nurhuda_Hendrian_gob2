package main

import "fmt"

func main() {

	fmt.Print("\n------------------------\n")
	var person map[string]string
	person = map[string]string{}
	person["name"] = "mnurhuda"
	person["age"] = "23"
	person["address"] = "Sleman, D.I.Yogyakarta"

	//Atau
	/*erson = map[string]string{
		person["name"]:    "mnurhudah",
		person["age"]:     "23",
		person["address"]: "Sleman, D.I.Yogyakarta",
	}*/

	fmt.Println("name : ", person["name"])
	fmt.Println("age : ", person["age"])
	fmt.Println("Address : ", person["address"])

	//Map x Looping
	fmt.Print("\n------------------------\n")
	var persoon = map[string]string{
		"name":    "mnurhudah",
		"age":     "23",
		"address": "Sleman, D.I.Yogyakarta",
	}

	for key, value := range persoon {
		fmt.Println(key, ": ", value)
	}

	//delete map
	fmt.Print("\n------------------------\n")
	var perrson = map[string]string{
		"name":    "mnurhudah",
		"age":     "23",
		"address": "Sleman, D.I.Yogyakarta",
	}

	fmt.Println("Before : ", perrson)

	delete(perrson, "address")
	fmt.Println("After : ", perrson)

	//Map (Detecting a value)
	fmt.Print("\n------------------------\n")
	var peerson = map[string]string{
		"name":    "mnurhudah",
		"age":     "23",
		"address": "Sleman, D.I.Yogyakarta",
	}
	value, exist := peerson["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value doesn't exist")
	}

	delete(peerson, "age")
	value, exist = peerson["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value has been deleted")
	}

	// Map (Combining slice with map)
	fmt.Print("\n------------------------\n")
	var people = []map[string]string{
		{"name": "Ariel", "age": "23"},
		{"name": "Asmiel", "age": "23"},
		{"name": "Fasiel", "age": "23"},
	}

	for i, person := range people {
		fmt.Printf("index: %d, name:%s, age: %s \n", i, person["name"], person["age"])
	}
}
