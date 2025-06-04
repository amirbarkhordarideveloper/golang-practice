package main

import "fmt"

func main() {
	var nnumber int
	fmt.Print("How many numbers do you want to enter:")
	fmt.Scan(&nnumber)

	var sum float64 = 0

	for i := 0; i < nnumber; i++ {
		var num float64
		fmt.Printf("Enter the number %d :", (i + 1))
		fmt.Scan(&num)
		sum += num
	}
	avrege := sum / float64(nnumber)
	fmt.Printf("avrege is %.2f\n", avrege)

}
