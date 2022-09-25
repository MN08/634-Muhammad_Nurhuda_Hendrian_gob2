package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	// FUNCTION
	fmt.Print("\n----------function-----------\n")
	biodata("mnh", 23)

	//jika parameter betipe data sama hanya perlu satu kali di deklarasikan tipedatanya
	welcome("mnurhuda", "good Morning")

	// Function (Return)
	names := []string{"alfaro", "manjiro"}
	printMsg := returnExample("Hi", names)

	fmt.Print("\n------------------------\n")
	fmt.Println(printMsg)

	// Function (Returning multiple values)
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Print("\n------------------------\n")
	fmt.Println("area : ", area)
	fmt.Println("circumference : ", circumference)

	//Function (Predefined return value)
	var diameters float64 = 15
	var areas, circumferences = calculates(diameters)
	fmt.Print("\n------------------------\n")
	fmt.Println("areas : ", areas)
	fmt.Println("circumferences : ", circumferences)
	fmt.Print("\n------------------------\n")

	// Function (Variadic function #1)
	nameLists := print("name1", "name2", "name3", "name4")
	fmt.Printf("%#v", nameLists)
	fmt.Print("\n------------------------\n")

	// Function (Variadic function #2)
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	result := sum(numberList...)
	fmt.Println(" Result : ", result)

	// Function (Variadic function #3)
	profile("jorgan", "jagung", "kacang", "mlinjo", "Bawang")

}

func biodata(name string, age int) {
	fmt.Print("\n------------------------\n")
	fmt.Printf("My name is %s, i am %d yo", name, age)
}

// jika parameter betipe data sama hanya perlu satu kali di deklarasikan tipedatanya
func welcome(name, greet string) {
	fmt.Print("\n------------------------\n")
	fmt.Printf("Hi %s, %s nice to meet you", name, greet)
}

// Function (Return)
func returnExample(msg string, names []string) string {
	joinStr := strings.Join(names, " ")
	fmt.Print("\n------------------------\n")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}

// Function (Returning multiple values)
func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

// Function (Predefined return value)
func calculates(d float64) (areas float64, circumferences float64) {
	areas = math.Pi * math.Pow(d/2, 2)
	circumferences = math.Pi * d

	return
}

// Function (Variadic function #1)
func print(name ...string) []map[string]string {
	var result []map[string]string

	for i, v := range name {
		key := fmt.Sprintf("name: %d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

// Function (Variadic function #2)
func sum(num ...int) int {
	total := 0

	for _, v := range num {
		total += v
	}
	return total
}

// Function (Variadic function #3)
func profile(namee string, favFood ...string) {
	mergeFavFood := strings.Join(favFood, ", ")

	fmt.Print("\n------------------------\n")
	fmt.Println("Hi, i am ", namee)
	fmt.Println("my fav food is :", mergeFavFood)

}
