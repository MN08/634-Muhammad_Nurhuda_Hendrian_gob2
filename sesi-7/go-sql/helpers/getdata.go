package helpers

import (
	"fmt"
)

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
