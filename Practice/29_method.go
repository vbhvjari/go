package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

type Cir struct {
	radius float64
}

// method name: Area() with receiver type
// receiver is r of type Rectangle
func (r Rectangle) Area() int {
	return r.length * r.width
}

// method name: Area() with receiver type
// receiver is c of type Cir
func (c Cir) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	r1 := Rectangle{length: 5, width: 10}
	c1 := Cir{radius: 12}

	fmt.Printf("\nArea of rectangle: %d", r1.Area())
	fmt.Printf("\nArea of circle: %f", c1.Area())
}
