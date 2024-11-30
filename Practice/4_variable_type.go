package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Vaibhav"
	var i int = 6
	var f float64 = 5.9
	var b bool = true

	// Type of variable using %T
	fmt.Printf("Name: %s and its type: %T", name, name)
	fmt.Printf("\nInteger: %d and its type: %T", i, i)
	fmt.Printf("\nFloat: %f and its type: %T", f, f)
	fmt.Printf("\nBoolean: %t and its type: %T", b, b)

	// Type of variable or value using reflect package
	// reflect.TypeOf(value)
	fmt.Printf("\nName: %s and its type: %v", name, reflect.TypeOf(name))

}
