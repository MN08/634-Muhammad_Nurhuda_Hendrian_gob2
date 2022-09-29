package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
)

func main() {
	//Simple Login
	fmt.Print("\n\n------------------------\n")
	var username string

	fmt.Print("Enter username : ")
	fmt.Scanln(&username)
	fmt.Print("Enter password : ")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		os.Exit(1)
	}

	if valid, err := loginValidation(username, password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func loginValidation(username string, password []byte) (string, error) {
	uname := "Anto"
	pass := "12345"
	p := string(password[:])
	if username != uname || p != pass {
		return "", errors.New("username / password not correct")
	} else {
		return "Login Success", nil
	}
}
