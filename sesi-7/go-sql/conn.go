package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to database successfully established...")

	//insert into database
	// CreateEmployees()

	//Update data
	// UpdateEmployees()

	//delete datq from database
	DeleteEmployees()

	//get data from database
	GetEmployees()

}

type Employee struct {
	employee_ID int
	full_name   string
	email       string
	age         int
	division    string
}

func CreateEmployees() {
	var employees = Employee{}

	sqlStatement := `INSERT INTO employees(full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *`

	err = db.QueryRow(sqlStatement, "Tugi", "tugi@golang.org", 34, "HRD").
		Scan(&employees.employee_ID, &employees.full_name, &employees.email, &employees.age, &employees.division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data: %+v\n", employees)
}

func GetEmployees() {
	var result = []Employee{}
	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.employee_ID, &employee.full_name, &employee.email, &employee.age, &employee.division)

		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}

	fmt.Println("Employee Data : ", result)
}

func UpdateEmployees() {
	sqlStatement := `UPDATE employees SET full_name = $2, email = $3, age = $4, division = $5 WHERE employee_id = $1;`
	res, err := db.Exec(sqlStatement, 3, "Tugino", "tuigno@golang.org", 34, "Finance")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Data amount : ", count)
}

func DeleteEmployees() {
	sqlStatement := `DELETE FROM employees WHERE employee_id = $1`

	res, err := db.Exec(sqlStatement, 2)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted Data amount:", count)
}
