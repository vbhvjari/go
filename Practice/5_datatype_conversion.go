package main

import (
	"fmt"
	"strconv"
)

func main() {
	// integer to float
	var i int = 10
	var f float64 = float64(i)
	fmt.Printf("Integer to float: %f", f)

	// Float to integer
	var f1 float64 = 5.6
	var i1 int = int(f1)
	fmt.Printf("\nFloat to integer: %d", i1)

	// integer to string using strconv package
	var i2 int = 10
	var s string = strconv.Itoa(i2)
	fmt.Printf("\n Integer to String : %q", s)

	// string to integer using strconv package
	var s1 string = "100abc"
	i, err := strconv.Atoi(s1)
	fmt.Printf("\n %v %T", i, i)
	fmt.Printf("\n %v %T", err, err)
}
