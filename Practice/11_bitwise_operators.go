package main

import "fmt"

func main() {
	var x, y int = 12, 15
	a := x & y
	fmt.Printf("Bitwise AND of %d and %d is %d", x, y, a)

	b := x | y
	fmt.Printf("\nBitwise OR of %d and %d is %d", x, y, b)

	c := x ^ y
	fmt.Printf("\nBitwise XOR of %d and %d is %d", x, y, c)

	d := x << 2
	fmt.Printf("\nBitwise left shift of %d by 2 is %d", x, d)

	e := x >> 2
	fmt.Printf("\nBitwise right shift of %d by 2 is %d", x, e)
}
