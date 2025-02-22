package main

import "fmt"

func sumNumbers(number ...int) int {
	sum := 0
	for _, val := range number {
		sum += val
	}
	return sum
}

func main() {
	fmt.Printf("Without any arguments: %d\n", sumNumbers())
	fmt.Printf("With 1 argument: %d\n", sumNumbers(5))
	fmt.Printf("With two arguments: %d\n", sumNumbers(5, 10))
}
