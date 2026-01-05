package main

import "fmt"

type Movie struct {
	name   string
	rating float32
}

func main() {
	m := Movie{"abc", 3.5}
	fmt.Printf("\nStruct m: %+v", m) // Output -> Struct m: {name:abc rating:3.5}
	fmt.Printf("\nStruct m: %v", m)  // Output -> Struct m: {abc 3.5}

}
