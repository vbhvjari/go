package main

import "fmt"

func modify(s []int) {
	s[0] = 100
}

func main() {
	s1 := []int{10, 20, 30}
	fmt.Printf("\nSlice s1: %v", s1)
	modify(s1)
	fmt.Printf("\nAfter modify call, slice s1: %v", s1)
}
