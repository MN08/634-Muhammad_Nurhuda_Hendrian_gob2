package helpers

import (
	"database/sql"
	"fmt"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	employee_ID int
	full_name   string
	email       string
	age         int
	division    string
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `INSERT INTO employee(full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *`

	err = db.QueryRow(sqlStatement, "Tugio", "tugio@golang.org", 34, "Human Resource Manager").
		Scan(&employee.employee_ID, &employee.full_name, &employee.email, &employee.age, &employee.division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data: %+v\n", employee)
}
