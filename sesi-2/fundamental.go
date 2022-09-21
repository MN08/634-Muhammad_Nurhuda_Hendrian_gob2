package main

import "fmt"

func main() {

	//conditional in go
	//if, if else, if else if
	fmt.Print("IF, IF ELSE, IF ELSE IF")
	fmt.Print("\n------------------------\n")

	name := "Anton"
	eat := false
	sleep := true

	if eat {
		fmt.Print(name, " sedang makan \n")
	}

	if eat {
	} else {
		fmt.Print(name, " tidak sedang makan \n")
	}

	if eat && !sleep {
		fmt.Print(name, " sedang makan")
	} else if !eat && sleep {
		fmt.Print(name, " sedang tidur")
	} else {
		fmt.Print(name, " tidak sedang makan ataupu  tidur")
	}

	//Switch Case
	fmt.Print("\n\n------------------------\n")
	fmt.Print("swtich case")
	fmt.Print("\n------------------------\n")
	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
	// lebih dari untuk 1 case
	fmt.Print("\n------------------------\n")
	var points = 6

	switch points {
	case 10:
		fmt.Println("perfect")
	case 11, 12, 13, 14:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// switch case fallthrough
	fmt.Print("\n------------------------\n")
	var pooint = 6

	switch {
	case pooint == 8:
		fmt.Println("perfect")
	case (pooint < 8) && (pooint > 3):
		fmt.Println("awesome")
		fallthrough
	case pooint < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//Nested Condition
	fmt.Print("\n------------------------\n")
	fmt.Print("Nested Condition")
	fmt.Print("\n------------------------\n")
	var pointt = 10

	if pointt > 7 {
		switch pointt {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointt == 5 {
			fmt.Println("not bad")
		} else if pointt == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if pointt == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
