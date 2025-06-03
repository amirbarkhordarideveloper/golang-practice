package main

import (
	"fmt"
	"strconv"
	"strings"

	"testgo/info"
)

func main() {
	printbanner()

	weight := getUserMetcic(info.WeightPrompt)
	height := getUserMetcic(info.HeightPrompt)

	bmi := calculateBMI(weight, height)

	printBMI(bmi)
}

func printbanner() {
	fmt.Println(info.BmiTitle)
	fmt.Println(info.Seperator)
}

func getUserMetcic(promptText string) (value float64) {
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	value, _ = strconv.ParseFloat(input, 64)

	return
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

func printBMI(bmi float64) {
	fmt.Printf("your bmi : %.2f", bmi)
}
