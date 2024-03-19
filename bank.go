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
	wantsDepositMoney := choice == 2
	wantsWithdrawMoney := choice == 3
	// wantsExit := choice == 4

	if wantsCheckBalance {
		fmt.Printf("Your balance is: %.2f\n", accountBalance)
	} else if wantsDepositMoney{
		fmt.Print("How much do you want to deposit? ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("Balance Updated! New balance: %.2f\n", accountBalance)
	} else if wantsWithdrawMoney{
		fmt.Print("How much do you want to withdraw? ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Printf("Balance Updated! New balance: %.2f\n",  accountBalance)
	} else {
		fmt.Print("Have a great day! Goodbye.")
	}
}