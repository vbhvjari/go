package main

import "fmt"

func main() {
	// Method-1
	var n int = 5

	// Method-2
	var s string
	s = "vaibhav"

	// Method-3 : Shorthand
	f := 45.5

	fmt.Printf("Value of n is %d", n)
	fmt.Printf("\nValue of s is %s", s)
	fmt.Printf("\nValue of f is %f", f)

	var i1, i2 int
	i1 = 5
	i2 = 10
	fmt.Printf("\nValue of i1 and i2 are %d and %d", i1, i2)

	var i3, i4 int = 15, 20
	fmt.Printf("\nValue of i3 and i4 are %d and %d", i3, i4)

	var i5, i6 int
	i5, i6 = 25, 30
	fmt.Printf("\nValue of i5 and i6 are %d and %d", i5, i6)

	var (
		in int     = 100
		sn string  = "n string"
		fn float32 = 100.10
		bn bool    = false
	)
	fmt.Printf("\nValue of in, sn, fn and bn are %d, %s, %f and %t respectively", in, sn, fn, bn)

	// Zero values
	var (
		iz int
		fz float32
		sz string
		bz bool
	)
	fmt.Printf("\nZero values of iz, fz, sz and bz are %d, %f, %s and %t respectively", iz, fz, sz, bz)

	const c1 int = 5
	const c2 = 10
	fmt.Printf("\nConstant values c1 and c2 are %d and %d", c1, c2)

}
