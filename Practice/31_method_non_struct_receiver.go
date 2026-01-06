package main

import "fmt"

/*
Error: cannot define new methods on non-local type int
func (a int) add(b int) {
}

func main() {

}
*/

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	n1 := myInt(5)
	n2 := myInt(10)
	sum := n1.add(n2)
	fmt.Printf("\n Addition: %d", sum)
}
