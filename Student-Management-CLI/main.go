package main

import "fmt"

func main() {
	students := make(map[string]int)

	var name1 string
	var name2 string
	var name3 string

	var age1 int
	var age2 int
	var age3 int

	var studentId1 int
	var studentId2 int
	var studentId3 int

	fmt.Println("Hello World")
	fmt.Print("enter th name of student number-1:")
	fmt.Scan(&name1)
	fmt.Print("enter the age of student number-1:")
	fmt.Scan(&age1)
	fmt.Print("enter the studentID of student number-1:")
	fmt.Scan(&studentId1)

	fmt.Print("enter th name of student number-2:")
	fmt.Scan(&name2)
	fmt.Print("enter the age of student number-2:")
	fmt.Scan(&age2)
	fmt.Print("enter the studentID of student number-2:")
	fmt.Scan(&studentId2)

	fmt.Print("enter th name of student number-3:")
	fmt.Scan(&name3)
	fmt.Print("enter the age of student number-3:")
	fmt.Scan(&age3)
	fmt.Print("enter the studentID of student number-3:")
	fmt.Scan(&studentId3)

	students[name1] = studentId1
	students[name2] = studentId2
	students[name3] = studentId3

	fmt.Printf("student number1\n: name:%v,age:%v,studentId:%v\n", name1, age1, studentId1)
	fmt.Printf("student number2\n: name:%v,age:%v,sutudentId:%v\n", name2, age2, studentId2)
	fmt.Printf("student number3: name:%v,age:%v,studentId:%v\n", name3, age3, studentId3)

	var searchstudent string
	fmt.Print("enter the student name for search:")
	fmt.Scan(&searchstudent)
	id, ok := students[searchstudent]
	if ok {
		fmt.Println("the student found: studentId:", id)
	} else {
		fmt.Println("the student not found")
	}
}
