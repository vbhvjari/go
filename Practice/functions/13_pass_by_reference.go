package main

import "fmt"

func modify(x *int) {
	*x = 100
}

func main() {
	n := 10
	fmt.Printf("\nValue of n: %v", n)
	modify(&n)
	fmt.Printf("\nValue of n after modify function: %v", n)
}
