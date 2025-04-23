package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	err := os.WriteFile(accountBalanceFile, []byte(fmt.Sprint(balance)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getBalanceFromFile() float64 {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}
