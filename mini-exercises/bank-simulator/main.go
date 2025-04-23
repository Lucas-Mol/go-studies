package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go Bank!")

	for {
		printInitialOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			printCurrentBalanceMessage()
		case 2:
			err := deposit()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			printUpdatedBalanceMessage()
		case 3:
			err := withdraw()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			printUpdatedBalanceMessage()
		default:
			exitGoodbye()
			return
		}
	}
}
