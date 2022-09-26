package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) perimeter() float64 {
	return (r.height * r.width) * 2
}

func main() {

	// Interface
	fmt.Print("\n------------------------\n")
	var c1 Shape = Circle{radius: 5}
	var r1 Shape = Rectangle{width: 3, height: 8}

	fmt.Printf("Type of C1 = %T\n", c1)
	fmt.Printf("Type of R1 = %T\rt", r1)

	fmt.Print("\n------------------------\n")
	fmt.Println("Circle Area : ", c1.area())
	fmt.Println("Circle Perimeter : ", c1.perimeter())
	fmt.Print("\n------------------------\n")
	fmt.Println("Rectangle Area : ", r1.area())
	fmt.Println("Rectangle Perimeter : ", r1.perimeter())

	//interface as functin parameter
	fmt.Print("\n\n------------------------\n")
	print("circle", c1)
	print("rectanlge", r1)

	// Interface (Type assertion)
	fmt.Print("\n\n------------------------\n")
	value, ok := c1.(Circle)

	if ok == true {
		fmt.Printf("Circle Value : %+v\n", value)
		fmt.Printf("Circle Volume : %f", value.volume())
	}
}

func print(t string, s Shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

// Interface (Type assertion)
func (c Circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
