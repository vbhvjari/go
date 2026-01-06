package main

import "fmt"

type Dog struct {
	name string
}

// Simple function
func bark(d Dog) {
	fmt.Printf("\nUsing function: %v", d.name)
}

// Method
func (d Dog) bark() {
	fmt.Printf("\nUsing Method: %v", d.name)
}

func main() {
	d := Dog{"tomy"}

	// Call function
	bark(d)

	// Call method
	d.bark()

}
