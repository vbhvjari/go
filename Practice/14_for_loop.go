package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	var a int = 1
	for a < 5 {
		fmt.Printf("Value of a: %d\n", a)
		a++
	}

	b := 1
	for b <= 5 {
		if b == 4 {
			break
		}
		fmt.Printf("Value of b: %d\n", b)
		b++
	}

	c := 1
	for c <= 5 {
		if c == 4 {
			c++
			continue
		}
		fmt.Printf("Value of c: %d\n", c)
		c++
	}

	// infinite loop
	/*
		for {
			fmt.Println("This loop will run forever.")
		}
	*/
}
