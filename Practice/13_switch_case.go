package main

import "fmt"

func main() {
	var a int = 10

	switch a {
	case 5:
		fmt.Println("a is 5")
	case 100, 200:
		fmt.Println("a is either 100 or 200")
	default:
		fmt.Println("a is neither 5 nor 100 nor 200")
	}

	var b int = 25

	switch b {
	case 5:
		fmt.Println("b is 5")
	case 25:
		fmt.Println("b is 25")
		fallthrough
	case 10:
		fmt.Println("b is 10")
		fallthrough
	default:
		fmt.Println("default")
	}

	var x, y int = 20, 10
	switch {
	case x+y == 30:
		fmt.Println("x + y is 30")
	case x-y == 10:
		fmt.Println("x - y is 10")
	default:
		fmt.Println("default")
	}
}
