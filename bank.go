package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("failed to find file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil{
		return 0, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main(){
	var accountBalance, err = getBalanceFromFile()

	if err!= nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		
		presentOptions()

		// Get user input

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1
		// wantsDepositMoney := choice == 2
		// wantsWithdrawMoney := choice == 3
		// wantsExit := choice == 4

		switch choice{
		case 1:
			fmt.Printf("Your balance is: %.2f\n", accountBalance)

		case 2:
			fmt.Print("How much do you want to deposit? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0{
					fmt.Println("Invalid amount. Please enter an amount great than 0.")
					continue
			}
			accountBalance += depositAmount
			fmt.Printf("Balance Updated! New balance: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Print("How much do you want to withdraw? ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance{
					fmt.Println("Invalid amount. Please enter an amount less than your balance.")
					continue
			}
			if withdrawAmount <= 0{
					fmt.Println("Invalid amount. Please enter an amount great than 0.")
					continue
			}
			accountBalance -= withdrawAmount
			fmt.Printf("Balance Updated! New balance: %.2f\n",  accountBalance)
			writeBalanceToFile(accountBalance)
		
		default:
			fmt.Println()
			fmt.Println("Have a great day! Goodbye.")
			fmt.Println("Thanks for choosing our bank.")
			return
			//break
		}
	}
}