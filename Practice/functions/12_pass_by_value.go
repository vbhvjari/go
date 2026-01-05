package main

import "fmt"

func modify(x int) { // x is parameter; when this function is called, copy of parameter is stored at separate memory location
	x = 100
}

func main() {
	n := 10
	fmt.Printf("\nValue of n: %v", n)
	modify(n) // Pass by value ; n is argument - we pass value of n when invoking function modify()
	fmt.Printf("\nValue of n after modify function: %v", n)
}
