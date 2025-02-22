package main

import "fmt"

func main() {
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Printf("Type of return value: %T\n", x)
	fmt.Printf("Return value of anonymous function: %d\n", x(20, 30))

	y := func(a int, b int) int {
		return a * b
	}(20, 30)

	fmt.Printf("Type of return value: %T\n", y)
	fmt.Printf("Return value of anonymous function: %d\n", y)
}
