package main

import (
	"fmt"
	"math"
) 

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	//var years float64 = 10

	fmt.Print("Please enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)
	//for only one words input - Scan() cannot take multiline input

	fmt.Print("Please enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	
	fmt.Print("Please enter the years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	var futureRealValue = futureValue * math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}