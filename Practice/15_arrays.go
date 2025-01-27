package main

import "fmt"

func main() {
	var grades [5]int
	fmt.Printf("Grades: %d\n", grades)

	var fruits [3]string
	fmt.Printf("Fruits: %s\n", fruits)

	var percentages [2]float64
	fmt.Printf("Percentages: %f\n", percentages)

	var result [2]bool
	fmt.Printf("Result: %t\n", result)

	// Array initialization
	var marks [3]int = [3]int{1, 2, 3}
	fmt.Printf("Marks: %d\n", marks)

	// shoryhand
	num := [3]int{1, 2, 3}
	fmt.Printf("Num: %d\n", num)

	// Partial initialization
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
