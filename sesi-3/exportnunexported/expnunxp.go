package main

import (
	"fmt"
	"helper"
)

func main() {

	fmt.Print("\n------------------------\n")

	helper.Greet()
	helper.greet()

	var person = helper.Person{}
	person.invokeGreet()
}
