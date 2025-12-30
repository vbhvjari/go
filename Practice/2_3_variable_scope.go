package main

import "fmt"

var g string = "global"

// fmt.Printf("Accessing global variable outside main function: %s", g)
// Error: syntax error: non-declaration statement outside function body
func main() {
	var l1 string = "local1"
	{
		var l2 string = "local2"
		fmt.Printf("Acesssing local variable within its scope: %s", l2)
	}
	fmt.Printf("\nAccess global variable inside main function: %s", g)
	fmt.Printf("\nAccessing local variable: %s", l1)

	// fmt.Printf("\nAccessing local variable outside its scope: %s", l2)
	// Error:  undefined: l2
}
