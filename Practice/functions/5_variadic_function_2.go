package main

import "fmt"

func printDetails(student string, subjects ...string) {
	fmt.Printf("Student name is : %s\n", student)
	fmt.Printf("Subjects of student are: \n")

	for _, val := range subjects {
		fmt.Printf("Subject name is: %s\n", val)
	}
}

func main() {

	printDetails("Vaibhav")
	printDetails("Vaibhav", "Maths")
	printDetails("Vaibhav", "Maths", "Physics")
}
