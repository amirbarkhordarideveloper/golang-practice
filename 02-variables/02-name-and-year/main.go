package main

import "fmt"

func main() {
	var firstname string = "Amir"
	lastname := "barkhordari"

	fmt.Println("my first name is:", firstname)
	fmt.Println("my last name is:", lastname)

	var Currentyear int = 2025
	brithyear := 2000

	age := Currentyear - brithyear
	fmt.Println("my age:", age)

	Currentyear = Currentyear + 1
	fmt.Println("next year is:", Currentyear)

}
