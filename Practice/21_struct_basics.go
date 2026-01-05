package main

import "fmt"

type Student struct {
	name   string
	rollno int
	marks  []int
	grades map[string]int
}

func main() {
	// Method-1 - Struct initialisation
	var s Student
	fmt.Printf("\nStruct: %v", s)
	fmt.Printf("\nStruct with special format: %+v", s)

	// Method-2 - Struct initialisation
	st := new(Student)
	fmt.Printf("\nStruct using new initialisation: %+v", st)

	// Method-3 - Struct initialisation - shorthand
	ss := Student{
		name:   "Vaibhav",
		rollno: 5,
		marks:  []int{10, 20, 30},
		grades: map[string]int{"A": 65, "F": 70},
	}
	fmt.Printf("\nStruct using shorthand initialisation: %+v", ss)

	// Method-4
	ss2 := Student{"Jariwala", 10, []int{50}, map[string]int{"K": 75}}
	fmt.Printf("\nStruct: %+v", ss2)

	var s2 Student
	s2.name = "Vbhv"
	s2.rollno = 15

	fmt.Printf("\nStruct instance s2: %+v", s2)
	fmt.Printf("\nRoll no from Struct instance s2: %+v", s2.rollno)

}
