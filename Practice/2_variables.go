package main

import "fmt"

// Global variables
var lan string = "Go"

// fmt.Printf("Access global variable outside main function: %s", lan) // Error: syntax error: non-declaration statement outside function body

func main() {
	var name string = "Vaibhav"
	var exp int = 7
	var male bool = true
	var height float64 = 5.9
	fmt.Print("Name : ", name, "\nExperience : ", exp, "\nMale :", male, "\nHeight : ", height)

	// Using Printf
	fmt.Printf("Name: %s", name)

	var my_name string = "Vaibhav"
	var i int = 80
	fmt.Printf("Hello %v, you scored %d/100", my_name, i)

	/* Format Specifiers
	%v	--> default format
	%d	--> decimal integers
	%f 	--> floating numbers
	%.2f --> floating numbers upto 2 decimal places
	%s	--> Plain string
	%c	--> character
	%q	--> quoted character/string
	%t	--> boolean - true or false
	%T	--> type of the value
	%v -> value in default format
	%T -> data type of the value
	*/

	// Short hand declaration

	// Multiple variables of same type
	var myname, surname string = "Vaibhav", "Jariwala"
	fmt.Print("\n")
	fmt.Printf("Full name is: %s %s", myname, surname)

	// Multiple variables of different type
	var (
		s  string = "Surat"
		in int    = 5
	)
	fmt.Print("\n")
	fmt.Printf("City: %s, Number: %d", s, in)

	// short variable declaration
	city := "Surat"
	fmt.Printf("\nCity: %s", city)

	// outer block
	outer_city := "Surat"

	{
		inner_city := "Udhana"
		fmt.Printf("\n From inner block -> inner_city: %s", inner_city)
		fmt.Printf("\n From inner block --> outer_city: %s", outer_city)
	}

	fmt.Printf("\n From outer block --> outer_city: %s", outer_city)
	// fmt.Printf("\n From outer block --> inner_city: %s", inner_city) // Error: undefined: inner_city

	// Access global variable
	fmt.Printf("\n Access global variable from main function: %s", lan)

	// Zero values
	var (
		sz string
		iz int
		fz float64
		bz bool
	)

	fmt.Printf("\n Zero values: String: %s, Integer: %d, Float: %f, Boolean: %t", sz, iz, fz, bz)
}
