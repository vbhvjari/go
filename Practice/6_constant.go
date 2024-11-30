package main

import "fmt"

func main() {

	//Typed constant
	const age int = 29
	const name string = "Vaibhav"

	fmt.Printf("Typed constants: %d, %s", age, name)

	//Untyped constant
	const a = 30
	const s = "Surat"

	fmt.Printf("\nUntyped constants: %d, %s", a, s)

}
