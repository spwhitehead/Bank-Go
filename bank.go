package main

import "fmt"

func main(){
	var accountBalance = 1000.00


	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	// Get user input

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	wantsCheckBalance := choice == 1
	// wantsDepositMoney := choice == 2
	// wantsWithdrawMoney := choice == 3
	// wantsExit := choice == 4

	if wantsCheckBalance {
		fmt.Println("Your balance is:", accountBalance)
	}

	fmt.Println("Your choice:", choice)
}