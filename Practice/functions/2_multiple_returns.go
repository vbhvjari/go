package main

import "fmt"

func operations(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff // called terminating statement
}

func main() {
	addition, difference := operations(10, 5)
	fmt.Printf("Addtion: %d and Subtraction: %d\n", addition, difference)

}
