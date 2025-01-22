package main

import "fmt"

func main() {
	var a int = 10

	if a < 50 {
		fmt.Printf("Condition is true\n")
	}

	var s string = "Vaibhav"
	if s == "vaibhav" {
		fmt.Printf("Condition for checking string is true\n")
	} else {
		fmt.Printf("Condition for checking string is false\n")
	}

	var f string = "grape"
	if f == "apple" {
		fmt.Printf("I like apples\n")
	} else if f == "orange" {
		fmt.Printf("I like oranges\n")
	} else {
		fmt.Printf("I like grapes\n")
	}
}
