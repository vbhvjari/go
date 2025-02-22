package main

import "fmt"

func greetings(st string) {
	fmt.Printf("Hey there, %s\n", st)
}

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	greetings("Vaibhav")
	output := addNumbers(6, 7)
	fmt.Printf("The outpus of addition is %d\n", output)
}
