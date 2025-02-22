package main

import "fmt"

func printName(s string) {
	fmt.Println(s)
}

func printRollNo(n int) {
	fmt.Println(n)
}

func printAddress(a string) {
	fmt.Println(a)
}

func main() {
	printName("Vaibhav")
	defer printRollNo(5)
	printAddress("Surat")
}
