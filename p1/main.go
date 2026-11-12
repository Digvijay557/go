package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Age    int
	Marks  float64
}

var max int
var c = -1

func add(student *[]Student, name string, age int, marks float64) {
	if len(*student) >= max {
		fmt.Println("Student panel is full!")
		return
	}

	*student = append(*student, Student{
		RollNo: len(*student) - 1 - c,
		Name:   name,
		Age:    age,
		Marks:  marks,
	})
	c = len(*student) - 1
}

func list(student *[]Student) {
	for i := 0; i < len(*student); i++ {
		fmt.Printf("RollNo: %d\nName: %s\nAge: %d\nMarks: %f\n", (*student)[i].RollNo+1, (*student)[i].Name, (*student)[i].Age, (*student)[i].Marks)
	}
}

func search(student *[]Student) int {
	var roll int
	fmt.Println("Enter the Roll no to Search:")
	_, err := fmt.Scan(&roll)
	if err != nil {
		fmt.Println("Invalid roll number")
		return -1
	}
	if roll < 1 || roll > len(*student) {
		fmt.Println("Student not found")
		return -1
	}

	stud := (*student)[roll-1]
	fmt.Printf("RollNo: %d\nName: %s\nAge: %d\nMarks: %f\n", stud.RollNo+1, stud.Name, stud.Age, stud.Marks)
	return roll
}

func update(student *[]Student) {
	roll := search(student)
	if roll == -1 {
		return
	}
	if roll < 1 || roll > len(*student) {
		fmt.Println("Invalid RollNo")
		return
	}

	stud := &(*student)[roll-1]
	var choice int

	for {
		fmt.Println("What do you Want to update?")
		fmt.Println("1.Name 2.Age 3.Marks 0.Exit")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid choice")
			return
		}

		switch choice {
		case 0:
			fmt.Println("Exit!!")
			return
		case 1:
			fmt.Println("Enter the new Name:")
			_, err = fmt.Scan(&stud.Name)
			if err != nil {
				fmt.Println("Invalid name")
			}
		case 2:
			fmt.Println("Enter the new Age:")
			_, err = fmt.Scan(&stud.Age)
			if err != nil {
				fmt.Println("Invalid age")
			}
		case 3:
			fmt.Println("Enter the new Marks:")
			_, err = fmt.Scan(&stud.Marks)
			if err != nil {
				fmt.Println("Invalid marks")
			}
		default:
			fmt.Println("Invalid option")
		}
	}
}

func delete(student *[]Student) {
	var roll int
	fmt.Println("Enter the roll no of student you want to delete:")
	if _, err := fmt.Scan(&roll); err != nil || roll < 1 || roll > len(*student) {
		fmt.Println("Invalid roll number")
		return
	}

	stud := &(*student)[roll-1]
	fmt.Printf("RollNo: %d\nName: %s\nAge: %d\nMarks: %f\n", stud.RollNo+1, stud.Name, stud.Age, stud.Marks)
	fmt.Println("Are you sure you want to delete [y/n]?")
	var confirmation string
	if _, err := fmt.Scan(&confirmation); err != nil || confirmation != "y" {
		return
	}

	stud.Name = ""
	stud.RollNo = -1
	stud.Marks = -1.0
	stud.Age = -1
}

func main() {
	fmt.Printf("Enter max number of students:")
	_, err := fmt.Scan(&max)
	if err != nil {
		fmt.Println("Invalid maximum number")
		return
	}

	students := make([]Student, 0, max)
	for {
		var choice int
		fmt.Println("Student Management System")
		fmt.Println("1.Add 2.List 3.Search 4.Update 5.Delete 0.Exit:")
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid choice")
			return
		}
		switch choice {
		case 0:
			fmt.Println("Exit")
			return
		case 1:
			fmt.Println("Add")
			var name string
			var age int
			var marks float64
			fmt.Printf("Enter Name:")
			fmt.Scan(&name)
			fmt.Printf("Enter Age:")
			fmt.Scan(&age)
			fmt.Printf("Enter Marks:")
			fmt.Scan(&marks)
			add(&students, name, age, marks)
		case 2:
			fmt.Println("List")
			list(&students)
		case 3:
			fmt.Println("Search")
			search(&students)
		case 4:
			fmt.Println("Update!!")
			update(&students)
		case 5:
			fmt.Println("Delete")
			delete(&students)
		}

	}
}
