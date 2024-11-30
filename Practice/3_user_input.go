package main

import "fmt"

func main() {

	// Single Input
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s", name)

	// Multiple Inputs
	var myname string
	var age int
	fmt.Print("\nEnter name and Age: ")
	fmt.Scanf("%s %d", &myname, &age)
	fmt.Printf("Name: %s, Age: %d", myname, age)

	/* Return values of Scanf
	count -> number of items successfully scanned
	err -> error if any
	*/

	var tech string
	fmt.Print("\nEnter your favourite technology: ")
	count, err := fmt.Scanf("%s", &tech)
	fmt.Printf("\nCount: %d, Error: %v", count, err)

}
