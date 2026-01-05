package main

import "fmt"

func modify(m map[string]int) {
	m["K"] = 75
}

func main() {
	m1 := make(map[string]int)
	m1["A"] = 65
	m1["F"] = 70
	fmt.Printf("\nMap m1: %v", m1)
	modify(m1)
	fmt.Printf("\nMap m1 after modify call: %v", m1)

}
