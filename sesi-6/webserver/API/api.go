package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "John", Age: 31, Division: "IT"},
	{ID: 2, Name: "Johny", Age: 41, Division: "Finance"},
	{ID: 3, Name: "Johnru", Age: 44, Division: "Sales"},
	{ID: 4, Name: "Johnson", Age: 34, Division: "Security"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployees)

	fmt.Println("Application listening on Port : ", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		//send data to html
		// tpl, err := template.ParseFiles("index.html")

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// tpl.Execute(w, employees)
		return
	}

	http.Error(w, "invalid Method", http.StatusBadRequest)
}

func createEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return

	}
	http.Error(w, "invalid method", http.StatusBadRequest)
}
