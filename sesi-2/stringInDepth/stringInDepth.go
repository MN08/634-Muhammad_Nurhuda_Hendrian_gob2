package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Strings In Depth
	// Looping Over String (byte-by-byte)
	fmt.Print("\n------------------------\n")
	word := "Iya"

	for i := 0; i < len(word); i++ {
		fmt.Printf("index: %d, byte:%d\n", i, word[i])
	}
	//sebaliknya
	fmt.Print("\n------------------------\n")
	var words []byte = []byte{73, 121, 97}
	fmt.Printf("Word : %s", string(words))

	// Looping Over String (rune-by-rune)
	fmt.Print("\n------------------------\n")
	str1 := "makan"
	str2 := "mânca"
	fmt.Printf("str1 byte length : %d\n", len(str1))
	fmt.Printf("str2 byte length : %d", len(str2)) //berjumlah 6 karna value byte dari â ada dua byte

	// Looping Over String (rune-by-rune)
	fmt.Print("\n------------------------\n")
	fmt.Printf("str1 rune length : %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 rune length : %d", utf8.RuneCountInString(str2))

	// proses to rune
	fmt.Print("\n------------------------\n")
	for i, s := range str2 {
		fmt.Printf("index: %d, string: %s\n", i, string(s))
	}
}
