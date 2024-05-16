package main

import "fmt"

type Student struct {
	Name  string
	Class string
	Marks []int
}

func (s Student) AverageMarks() float64 {
	sum := 0
	for _, mark := range s.Marks {
		sum += mark
	}
	if len(s.Marks) == 0 {
		return 0
	}
	return float64(sum) / float64(len(s.Marks))
}

func (s Student) Details() string {
	return fmt.Sprintf("%s is in class %s and has an average mark of %.2f", s.Name, s.Class, s.AverageMarks())
}

type Describer interface {
	Details() string
}

func main() {
	student1 := Student{
		Name:  "Alice",
		Class: "10A",
		Marks: []int{88, 90, 92, 85, 90},
	}
	student2 := Student{
		Name:  "Bob",
		Class: "10B",
		Marks: []int{75, 80, 65, 88, 80},
	}

	printDetails(student1)
	printDetails(student2)
}

func printDetails(d Describer) {
	fmt.Println(d.Details())
}
