package main

import (
	"fmt"
	"log"
)

func main() {
	for {
		fmt.Println(`
=============================
Welcome to Profit Calculator!
=============================
		`)
		// Not necessary, could just return values from functions.
		// Just to practice and understand pointers better
		var revenue float64
		var expenses float64
		var taxRate float64

		getUserInput("Revenue", &revenue)
		getUserInput("Expenses", &expenses)
		getUserInput("Tax Rate", &taxRate)

		ebt, profit, ratio := extractEBTProfitRatio(revenue, expenses, taxRate)

		fmt.Println("EBT:", ebt)
		fmt.Println("Profit:", profit)
		fmt.Printf("Ratio Result: %.2f\n", ratio)
	}
}

// The 'filledVariable' which receives the input value NEED to be a pointer,
// so it can be written in same local on memory and so "persisted" above this function scope
func getUserInput(requiredValueName string, filledVariable *float64) {
	fmt.Print(requiredValueName, ": ")
	_, err := fmt.Scan(filledVariable)
	if err != nil {
		log.Fatal(err)
	}
}

func extractEBTProfitRatio(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
