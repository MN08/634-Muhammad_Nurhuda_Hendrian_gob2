package main

import "fmt"

func main() {

	// Channels (Channel operators)
	fmt.Print("\n------------------------\n")
	c := make(chan string)
	// mengirim data ke channel
	// c <- value
	//menerima data dari channel
	// result := <-c

	// Channels (Implementing channel)
	fmt.Print("\n------------------------\n")
	go introduce("Gokuu", c)
	go introduce("Gohan", c)
	go introduce("Goten", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	// Channels (Channel with anonymous function)
	fmt.Print("\n\n------------------------\n")
	ch := make(chan string)

	students := []string{"Gogeta", "Gotunk", "Vegoku"}
	for _, v := range students {
		go func(student string) {
			fmt.Println("Student ", student)
			result := fmt.Sprintf("hey, my name is %s", student)
			ch <- result
		}(v)
	}

	for i := 0; i < len(students); i++ {
		print(ch)
	}

	close(ch)

	//Channels (Channel directions)
	fmt.Print("\n\n------------------------\n")

}

// Channels (Implementing channel)
func introduce(character string, c chan string) {
	result := fmt.Sprintf("hey, i am %s", character)

	c <- result
}

// Channels (Channel with anonymous function)
func print(ch chan string) {
	fmt.Println(<-ch)
}
