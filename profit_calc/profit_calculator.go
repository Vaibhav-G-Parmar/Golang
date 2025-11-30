//ask user for revenue, expenses and tax rates

package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRates float64

	fmt.Print("Please enter the Revenue Amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter the Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Please enter the Tax Rates: ")
	fmt.Scan(&taxRates)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRates/100)
	ratio := earningsBeforeTax / profit

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}