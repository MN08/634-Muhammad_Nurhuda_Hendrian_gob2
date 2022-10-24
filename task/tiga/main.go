package main

import "fmt"

func IsPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {

		if word[i] != word[len(word)-i-1] {

			return false

		}

	}

	return true

}

func main() {

	var str string

	fmt.Scan(&str)

	result := isPalindrome(str)

	if result == true {

		fmt.Println("palindrome")

	} else {

		fmt.Println("not palindrome")

	}

}
