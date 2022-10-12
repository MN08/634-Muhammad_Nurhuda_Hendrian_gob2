package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// func main() {
// 	// Decoding JSON To Struct
// 	fmt.Print("\n-----------Decoding JSON To Struct-------------\n")
// 	var jsonString = `
// 		{
// 			"full_name": "ahmad subaedah",
// 			"email": "ahmad@mjd.online",
// 			"age": 23
// 		}
// 	`

// 	var result Employee
// 	var err = json.Unmarshal([]byte(jsonString), &result)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("full_name", result.FullName)
// 	fmt.Println("email", result.Email)
// 	fmt.Println("age", result.Age)

// }

// func main() {
// 	// Decoding JSON To Map
// 	fmt.Print("\n-----------Decoding JSON To MAP-------------\n")
// 	var jsonString = `
// 		{
// 			"full_name": "ahmad subaedah",
// 			"email": "ahmad@mjd.online",
// 			"age": 23
// 		}
// 	`

// 	var result map[string]interface{}
// 	var err = json.Unmarshal([]byte(jsonString), &result)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("full_name", result["full_name"])
// 	fmt.Println("email", result["email"])
// 	fmt.Println("age", result["age"])
// }

// func main() {
// 	// Decoding JSON To Empty Interface
// 	fmt.Print("\n-----------Decoding JSON To Empty Interface -------------\n")
// 	var jsonString = `
// 		{
// 			"full_name": "ahmad subaedah",
// 			"email": "ahmad@mjd.online",
// 			"age": 23
// 		}
// 	`
// 	var temp interface{}
// 	var err = json.Unmarshal([]byte(jsonString), &temp)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	var result = temp.(map[string]interface{})

// 	fmt.Println("full_name", result["full_name"])
// 	fmt.Println("email", result["email"])
// 	fmt.Println("age", result["age"])

// }
func main() {
	// Decoding JSON Array To Slice Of Struct
	fmt.Print("\n-----------Decoding JSON Array To Slice Of Struct  -------------\n")
	var jsonString = `
		[
			{
				"full_name": "ahmad subaedah",
				"email": "ahmad@mjd.online",
				"age": 23
			},
			{
				"full_name": "subaedah ahmad",
				"email": "subaedah@mjd.online",
				"age": 23
			}
		]
	`

	var result []Employee
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf(" Index %d: %+v\n", i+1, v)
	}

}
