package main

import "fmt"

func main() {
	P := 3.14
	radius := 5.00

	Circumference := (P * radius) * 2
	fmt.Println("Circumference cricle:", Circumference)

	fmt.Printf("for a radius of %v , the circle ciricumference is %.2f", radius, Circumference)

}
