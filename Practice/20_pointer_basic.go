package main

import "fmt"

func main() {
	i := 10
	// Method-1 - Pointer Initialisation -  with explict data type
	var ptr_i *int = &i
	fmt.Printf("\nValue of i: %v and its type: %T", i, i)
	fmt.Printf("\nPointer variable value: %v and its type: %T", ptr_i, ptr_i)

	n := 20
	// Method-2 - Pointer Initialisation - without explicit data type
	var ptr_n = &n
	fmt.Printf("\nValue of n: %v and its type: %T", n, n)
	fmt.Printf("\nPointer variable value: %v and its type: %T", ptr_n, ptr_n)

	m := 30
	// Method-3 - Pointer Initialisation - Shorthand
	ptr_m := &m
	fmt.Printf("\nValue of m: %v and its type: %T", m, m)
	fmt.Printf("\nPointer variable value: %v and its type: %T", ptr_m, ptr_m)

	// Dereference the pointer
	x := *ptr_i
	fmt.Printf("\nValue of x: %v and its type: %T", x, x)

	// Change value using dereference operator
	*ptr_i = 100
	fmt.Printf("\nValue of i: %v and its type: %T", i, i)

}
