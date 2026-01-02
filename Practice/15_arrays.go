package main

import "fmt"

func main() {
	var grades [5]int
	fmt.Printf("Grades: %d\n", grades) // Zero value without array initialization

	var fruits [3]string
	fmt.Printf("Fruits: %s\n", fruits) // Zero value without array initialization

	var percentages [2]float64
	fmt.Printf("Percentages: %f\n", percentages) // Zero value without array initialization

	var result [2]bool
	fmt.Printf("Result: %t\n", result) // Zero value without array initialization

	// Array initialization
	var marks [3]int = [3]int{1, 2, 3}
	fmt.Printf("Marks: %d\n", marks)

	// Initialize with specific index
	var a1 [3]int = [3]int{
		0: 5,
		1: 10,
	}
	fmt.Printf("\nArray a1: %v", a1)

	// shorthand
	num := [3]int{1, 2, 3}
	fmt.Printf("Num: %d\n", num)

	// Partial initialization - rest zero values are considered
	numbers := [5]int{1, 2, 3}
	fmt.Printf("Numbers: %d\n", numbers)

	// elipsis
	ranks := [...]int{1, 2, 3}
	fmt.Printf("Ranks: %d\n", ranks)

	// length of array
	fruit := [2]string{"apple", "strawberry"}
	fmt.Printf("Length of fruits: %d\n", len(fruit))

	// Accessing elements
	fmt.Printf("Fruit at index 0: %s\n", fruit[0])

	// array is mutable
	fruit[0] = "banana"
	fmt.Printf("Fruit at index 0: %s\n", fruit[0])

	// loop
	nums := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("elements of nums array at index %d: %d\n", i, nums[i])
	}

	// using range
	for index, element := range nums {
		fmt.Printf("Using ranges-elements of nums array at index %d: %d\n", index, element)
	}

	// multidimentioonal array
	// 2D array
	matrix := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("Matrix: %d\n", matrix)
	fmt.Printf("first element of matrix: %d\n", matrix[0])
	fmt.Printf("second element of first element of matrix: %d\n", matrix[0][1])

}
