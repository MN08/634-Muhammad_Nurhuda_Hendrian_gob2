package main

import "fmt"

func main() {

	// Empty Interface
	fmt.Print("\n------------------------\n")

	var randomValue interface{}
	_ = randomValue

	randomValue = "jalan Kaki"
	randomValue = 20
	randomValue = true
	randomValue = []string{"My Name", " Is"}

	//Empty interface (Type assertion)
	fmt.Print("\n\n------------------------\n")
	var v interface{}

	v = 20
	// v = v * 9  error karna interface tuidak bisa dikali in

	if values, oke := v.(int); ok == true {
		v = values * 9
	}

	//Empty interface (Map & Slice with Empty Interface)
	fmt.Print("\n\n------------------------\n")
	rs := []interface{}{1, "ariel", true, 2, "noah", true}

	rm := map[string]interface{
		"name": "aril"
		"status": true
		"number": 3
	}

	_,_ =rs,rm
}
