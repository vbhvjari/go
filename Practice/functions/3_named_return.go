package main

import "fmt"

func operations(a int, b int) (sum int, diff int) {
	sum = a + b // no need to use short hand operator since sum is already defined in above return - function defintion
	diff = a - b
	return
}

func main() {
	addition, subtraction := operations(10, 5)
	fmt.Printf("Addition: %d and Subtraction: %d\n", addition, subtraction)

}
