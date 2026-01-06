package main

import "fmt"

type Shape interface {
	area() float64
}

type Circlee struct {
	radius float64
}

func (c Circlee) area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Printf("\nArea: %v", s.area())
}

func main() {
	c := Circlee{5}
	printArea(c)
}
