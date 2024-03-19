package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spwhitehead/Bank-Go/fileops"
)

const accountBalanceFile = "balance.txt"

func main(){
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err!= nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Located exclusively in:", randomdata.Country(randomdata.FullCountry))
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		
		default:
			fmt.Println()
			fmt.Println("Have a great day! Goodbye.")
			fmt.Println("Thanks for choosing our bank.")
			return
			//break
		}
	}
}