package main

import (
	"fmt"
	"strconv"
	"strings"

	"testgo/info"
)

func main() {
	fmt.Println(info.BmiTitle)
	fmt.Println(info.Seperator)

	fmt.Print(info.WeightPrompt)
	weightinput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightinput, _ := reader.ReadString('\n')

	weightinput = strings.TrimSpace(weightinput)
	heightinput = strings.TrimSpace(heightinput)

	weight, _ := strconv.ParseFloat(weightinput, 64)
	height, _ := strconv.ParseFloat(heightinput, 64)

	fmt.Println("this is your weight:", weight)
	fmt.Println("this is your height:", height)

	bmi := weight / (height * height)

	fmt.Printf("your bmi : %.2f", bmi)
}
