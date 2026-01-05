package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("\nValue of i: %v and its type: %T", i, i)

	// Address operator
	x := &i
	fmt.Printf("\nValue of x: %v and its type: %T", x, x)

	// Dereference operator
	y := *x
	fmt.Printf("\nValue of y: %v and its type: %T", y, y)
}
