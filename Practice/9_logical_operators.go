package main

import "fmt"

func main() {
	var a, b int = 10, 20
	fmt.Println((a < b) && (a < 50))
	fmt.Println((a > b) || a < 50)
	fmt.Println(!(a > b))
	fmt.Println(!(true))
	fmt.Println(!(false))

	var x, y bool = true, false
	fmt.Println(x && y)
	fmt.Println(x || y)
	fmt.Println(!x)
	fmt.Println(!y)
}
