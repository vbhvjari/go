package main

import "fmt"

func main() {
	var s1 string = "Vaibhav"
	var s2 string = "vaibhav"
	fmt.Print("Euqal comparision operator: ", s1 == s2)
	fmt.Print("\nNot equal comparision operator: ", s1 != s2)

	var a, b int = 10, 20
	fmt.Printf("\nEuqal to operator-Is a euqal to b:  %v", a == b)
	fmt.Printf("\nNot Euqal to operator-Is a not euqal to b:  %v", a != b)
	fmt.Printf("\nLess than euqal to operator--Is a less than euqal to b: %v", a <= b)
	fmt.Printf("\nGreater than operator-Is a greater than b:  %v", a > b)
	fmt.Printf("\nGreater than euqal to operator--Is a greater than euqal to b: %v", a >= b)
}
