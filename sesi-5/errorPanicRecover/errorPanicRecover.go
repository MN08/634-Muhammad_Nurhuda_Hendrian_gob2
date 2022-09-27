package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	//error
	fmt.Print("\n------------------------\n")
	var number int
	var err error

	number, err = strconv.Atoi("123HG")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	//Error (Custom error)
	fmt.Print("\n\n------------------------\n")
	var password string
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}

	//Recover
	defer catchError()
	//Panic
	fmt.Print("\n\n------------------------\n")
	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

// Error (Custom error)
func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password too short, min 5 characters")
	}
	return "valid password", nil
}

func catchError() {
	if r := recover(); r != nil {
		fmt.Println("Error Occured: ", r)
	} else {
		fmt.Println("All Goods")
	}
}
