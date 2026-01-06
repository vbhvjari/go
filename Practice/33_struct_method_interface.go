package main

import "fmt"

// Struct - holds data
type StudentId struct {
	name   string
	rollno int
	grade  int
}

// Method - actions related data
func (s StudentId) speakEnglish() {
	fmt.Printf("\nFrom StudentId method: Hello I am :%v", s.name)
}

// Method - actions related data
func (s StudentId) writeEmail() {
	fmt.Printf("\nFrom StudentId method: %v can write email", s.name)
}

type TeacherId struct {
	name    string
	subject string
}

func (t TeacherId) speakEnglish() {
	fmt.Printf("\nFrom TeacherId method: Hello I am :%v", t.name)
}

func (t TeacherId) writeEmail() {
	fmt.Printf("\nFrom TeacherId method: %v can write email", t.name)
}

// Interface
type Rule interface {
	speakEnglish()
	writeEmail()
}

func Program(r Rule) {
	r.speakEnglish()
	r.writeEmail()
}

func main() {
	s1 := StudentId{name: "sam", rollno: 1, grade: 1}
	t1 := TeacherId{name: "tam", subject: "english"}

	Program(s1)
	Program(t1)

}
