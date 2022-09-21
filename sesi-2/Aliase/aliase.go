package main

import "fmt"

func main() {
	//ALIASE
	fmt.Print("\n------------------------\n")
	var a uint8 = 10 //declare variable tipe data uint8
	var b byte       // byte adalah aliase dari tipe data uint8

	b = a // tidak error karna tipe data sama
	_ = b

	fmt.Println(a)
	fmt.Println(b)

	//aliase tipe data buatan sendiri

	fmt.Print("\n------------------------\n")
	//declare variable tipe data second
	// syntax type nama_alias = nama_tipe_data

	type second = uint
	var hour second = 3600
	fmt.Printf("hour type = %T \n", hour)

}
