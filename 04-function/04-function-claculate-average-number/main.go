package main

import "fmt"

func main() {

	nnumber := getnumberusers()

	sum := getNumberSum(nnumber)

	average := calculateaverege(sum, nnumber)
	fmt.Printf("averege is %.2f\n", average)

}

func getnumberusers() int {
	var nnumber int
	fmt.Print("How many numbers do you want to enter:")
	fmt.Scan(&nnumber)
	return nnumber
}

func getNumberSum(n int) float64 {
	var sum float64 = 0
	var num float64
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the number %d :", (i + 1))
		fmt.Scan(&num)
		sum += num
	}
	return sum

}
func calculateaverege(sum float64, nnumber int) float64 {
	return sum / float64(nnumber)
}
