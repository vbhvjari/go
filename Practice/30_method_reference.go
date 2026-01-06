package main

import "fmt"

type Employeee struct {
	name     string
	salary   int
	currency string
}

// displaySalary() is method of type Employeee
// e is receiver of type Employeee
func (e *Employeee) displayeSalary() {
	fmt.Printf("\nSalary of Persion %s is %d in %s format", e.name, e.salary, e.currency)
}

func main() {
	emp1 := Employeee{name: "sam", salary: 5000, currency: "$"}
	emp1.displayeSalary()
	emp1.salary = 6000
	emp1.displayeSalary()
}
