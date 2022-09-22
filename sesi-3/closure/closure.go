package main

import (
	"fmt"
	"strings"
)

type isOddNumb func(int) bool

func main() {

	// Closure(Declare closure in variable)
	fmt.Print("\n------------------------\n")
	var eventNumber = func(number ...int) []int {
		var result []int

		for _, v := range number {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	var number = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(eventNumber(number...))

	// Closure(IIFE)
	fmt.Print("\n------------------------\n")
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	//Closure (Closure as a return value)
	fmt.Print("\n------------------------\n")
	var studentsList = []string{"agus", "bagus", "cagus", "dagus", "eagus", "fagus", "gagus"}
	var find = findStudent(studentsList)
	fmt.Println(find("dagus"))

	//Closure (Callback)
	fmt.Print("\n------------------------\n")

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	var finds = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total Odd Number: ", finds)

}

// Closure (Closure as a return value)
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("No students found", s)
		}
		return fmt.Sprintf("we found %s at position %d", s, position+1)
	}
}

// Closure (Callback)
func findOddNumbers(numbers []int, callback isOddNumb) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
