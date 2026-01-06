package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

// displaySalary() is method of type Employee
// e is receiver of type Employee
func (e Employee) displaySalary() {
	fmt.Printf("\nSalary of Persion %s is %d in %s format", e.name, e.salary, e.currency)
}

func main() {
	emp1 := Employee{name: "sam", salary: 5000, currency: "$"}
	emp1.displaySalary()
}

/*
Same program only using function

package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

func displaySalary(e Employee) {
	fmt.Printf("\nSalary of Persion %s is %d in %s format", e.name, e.salary, e.currency)
}

func main() {
	emp1 := Employee{name: "sam", salary: 5000, currency: "$"}
	displaySalary(emp1)
}
*/
