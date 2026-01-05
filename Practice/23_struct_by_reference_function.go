package main

import "fmt"

type Circle struct {
	x int
	y int
	r float64
	a float64
}

func calArea(c *Circle) {
	const PI float64 = 3.14
	area := (PI * c.r * c.r)
	c.a = area
	fmt.Printf("\nValue of area from calArea: %+v", c.a)
}

func main() {
	c1 := Circle{x: 5, y: 5, r: 5, a: 0}
	fmt.Printf("\nstruct instance c1: %+v", c1)
	calArea(&c1) // Pass c1 as value
	fmt.Printf("\nStruct instance c1: %+v", c1)

}
