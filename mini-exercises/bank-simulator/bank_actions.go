package main

import (
	"errors"
	"fmt"
)

func deposit() error {
	balance := getBalanceFromFile()

	fmt.Print("Your deposit: ")
	var depositAmount float64
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		return errors.New("invalid amount. Must be greater than 0")
	}
	balance += depositAmount
	writeBalanceToFile(balance)
	return nil
}

func withdraw() error {
	balance := getBalanceFromFile()

	fmt.Print("Withdraw amount: ")
	var withdrawAmount float64
	fmt.Scan(&withdrawAmount)

	if withdrawAmount <= 0 {
		return errors.New("invalid amount. Must be greater than 0")
	}

	if withdrawAmount > balance {
		return errors.New("invalid amount. You can't withdraw more than you have")
	}

	balance -= withdrawAmount
	writeBalanceToFile(balance)
	return nil
}

func printCurrentBalanceMessage() {
	balance := getBalanceFromFile()
	fmt.Println("Your balance is", balance)
}

func printUpdatedBalanceMessage() {
	balance := getBalanceFromFile()
	fmt.Println("Balance updated! Your balance is", balance)
}
