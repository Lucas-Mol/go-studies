package main

import "fmt"

func printInitialOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func exitGoodbye() {
	fmt.Println("Goodbye!")
	fmt.Println("Thanks to choosing our bank")
}
