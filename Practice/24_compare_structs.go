package main

import "fmt"

type s struct {
	x int
}

func main() {
	c1 := s{x: 5}
	c2 := s{x: 10}

	if c1 == c2 {
		fmt.Printf("\nBoth structs are equal")
	} else {
		fmt.Printf("\nBoth structs are not equal")
	}
}
