package main

import (
	"fgd-golang/sesi-3/exportnunexported/helpers"
	"fmt"
)

func main() {

	fmt.Print("\n------------------------\n")

	helpers.Greet()
	// helpers.greet()

	var person = helpers.Person{}
	person.Invokegreet()
}
