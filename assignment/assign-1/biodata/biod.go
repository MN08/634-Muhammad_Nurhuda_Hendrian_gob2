package main

import (
	"errors"
	"fmt"
	"os"
)

// main layer
func mains() {
	// Give error if not available the second index (1) of args
	if len(os.Args) < 2 {
		fmt.Println("error : index student needed")
		return
	}
	// Get student args
	studentIndexArgs := os.Args[1]

	// Get student by index on other function / level
	searchedStudent, err := getStudentByIndex(studentIndexArgs)
	if err != nil {
		// show the error if from getStudentByIndex return error
		fmt.Println("error : ", err.Error())
		return
	}

	// Print the student
	fmt.Println(searchedStudent)

}

// domain layer and will call repo getStudentData
// the domain layer will handle logic of the data like search (if the search cannot be handled by repo layer)
func getStudentByIndex(index string) (*Student, error) {
	// return error if the index is empty
	if index == "" {
		return nil, errors.New("index is empty")
	}

	// get student data first
	studentData := getStudentData()

	// search the student using map key
	searchedStudent := studentData[index]

	// if searchcStudent is empty return error
	if searchedStudent == nil {
		return nil, errors.New("student not found")
	}

	return searchedStudent, nil
}

// layer repo
// construct data and use map with pointer data to easily get the null / exist data
// usually call the data from other data source like database etc
func getStudentData() map[string]*Student {
	return map[string]*Student{
		"1": {
			Index: 1,
			Name:  "Huda",
			Other: "Programmer",
		},
		"2": {
			Index: 2,
			Name:  "Wakhid",
			Other: "Chef",
		},
		"3": {
			Index: 3,
			Name:  "Pandu",
			Other: "Arlnod",
		},
	}
}

// Student entity layer
type Student struct {
	Index int
	Name  string
	Other string
}
