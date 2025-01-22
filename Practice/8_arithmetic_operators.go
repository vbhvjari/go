package main

import "fmt"

func main() {
	var s1, s2 string = "Vaibhav", "Jariwala"
	fmt.Printf("Addtion of two strings: %s and %s is %s", s1, s2, s1+s2)

	var a, b int = 20, 10
	fmt.Printf("\nAddtion of two numbers: %d and %d is %d", a, b, a+b)
	fmt.Printf("\nSubtration of two numbers: from %d minus %d is %d", a, b, a-b)
	fmt.Printf("\nMultiplication of two numbers: %d and %d is %d", a, b, a*b)
	fmt.Printf("\nDivision of two numbers: %d divided by %d is %d", a, b, a/b)
	fmt.Printf("\nModulus of two numbers: %d divided by %d is %d", a, b, a%b) // Returns remainder when left operand is divided by right operand
	a++                                                                       // Increment a by 1 by default
	fmt.Printf("\nIncrement of a: %d", a)
	b-- // Decrement b by 1 by default
	fmt.Printf("\nDecrement of b: %d", b)
}
