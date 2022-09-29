package main

import (
	"fmt"
	"time"
)

func main() {

	// Channels (Channel operators)
	fmt.Print("\n------------------------\n")
	c := make(chan string)
	// mengirim data ke channel
	// c <- value
	//menerima data dari channel
	// result := <-c

	// Channels (Implementing channel)
	fmt.Print("\n-----------Implementing channel-------------\n")
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
	fmt.Print("\n\n-----------Channel with anonymous function-------------\n")
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
	fmt.Print("\n\n-----------Channel directions-------------\n")
	cha := make(chan string)

	studentss := []string{"Gogeta", "Gotunk", "Vegoku"}

	for _, v := range studentss {
		go introduces(v, cha)
	}

	for i := 0; i < len(studentss); i++ {
		prints(cha)
	}

	close(cha)

	//Channels (Buffered vs unbuffered channel)
	fmt.Print("\n\n------------------------\n")
	fmt.Print("buffered channel\n")
	ch1 := make(chan int)

	go func(chenn chan int) {
		fmt.Println("func Goroutines start sending data into the channel")
		chenn <- 10
		fmt.Println("func Goroutines finished sending data into the channel")
	}(ch1)

	fmt.Println("go routine sleep 2 second")
	time.Sleep(time.Second * 2)

	fmt.Println("main go routine start reciving data")
	d := <-ch1
	fmt.Println("main go routine recieved data", d)

	close(ch1)
	time.Sleep(time.Second)

	fmt.Print("\n------------------------\n")
	fmt.Print("\n\nunbuffered channel\n")
	ch2 := make(chan int)
	go func(chenn2 chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func Goroutines %d start sending data into the channel\n", i)
			chenn2 <- i
			fmt.Printf("func Goroutines #%d finished sending data into the channel", i)
		}
		close(chenn2)
	}(ch2)

	fmt.Println("go routine sleep 2 second")
	time.Sleep(time.Second * 2)
	for v := range ch2 {
		fmt.Println("main go routine recieved data :", v)
	}

	// Channel (Select)
	fmt.Print("\n\n------------Select------------\n")
	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch3 <- "Wow!!!"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch4 <- "salut!!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-ch3:
			fmt.Println("Recived ", msg1)
		case msg2 := <-ch4:
			fmt.Println("Recived ", msg2)
		}
	}
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

// Channels (Channel directions)
func prints(cha chan string) {
	fmt.Println(<-cha)
}

func introduces(chara string, cha chan<- string) {
	result := fmt.Sprintf("hey,my name is %s", chara)
	cha <- result

}
