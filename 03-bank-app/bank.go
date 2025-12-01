package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// printMenu displays the available options to the user
func printMenu() {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

// readInput reads a line of input from the user
func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func main() {

	var accountBalance float64 = 0.00
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n##################################")
	fmt.Println("#    Welcome to the Go Bank!     #")
	fmt.Println("##################################")

	for {
		printMenu()
		fmt.Print("\nKindly enter your choice: ")
		choiceStr, err := readInput(reader)
		if err != nil {
			fmt.Println("\n### Error reading input, please try again. ###")
			continue
		}

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("\n### Invalid choice, please enter a number. ###")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("\n#################################")
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
			fmt.Println("#################################")
		case 2:
			fmt.Print("\nEnter amount to deposit: ")
			depositAmountStr, err := readInput(reader)
			if err != nil {
				fmt.Println("\n### Error reading input, please try again. ###")
				continue
			}
			depositAmount, err := strconv.ParseFloat(depositAmountStr, 64)
			if err != nil {
				fmt.Println("\n### Invalid amount. Please enter a number. ###")
				continue
			}

			if depositAmount <= 0 {
				fmt.Println("\n### Deposit amount must be positive. ###")
			} else {
				accountBalance += depositAmount
				fmt.Println("\n#################################")
				fmt.Printf("Successfully deposited $%.2f. \n  New balance is $%.2f\n", depositAmount, accountBalance)
				fmt.Println("#################################")
			}
		case 3:
			fmt.Print("\nEnter amount to withdraw: ")
			withdrawAmountStr, err := readInput(reader)
			if err != nil {
				fmt.Println("\n### Error reading input, please try again. ###")
				continue
			}
			withdrawAmount, err := strconv.ParseFloat(withdrawAmountStr, 64)
			if err != nil {
				fmt.Println("\n### Invalid amount. Please enter a number. ###")
				continue
			}

			if withdrawAmount <= 0 {
				fmt.Println("\n### Withdrawal amount must be positive. ###")
			} else if withdrawAmount > accountBalance {
				fmt.Println("\n### Insufficient funds for this withdrawal. ###")
				fmt.Printf("### Available balance is $%.2f ###\n", accountBalance)
			} else {
				accountBalance -= withdrawAmount

				fmt.Println("#################################")
				fmt.Printf("Successfully withdrew $%.2f. \n  New balance is $%.2f\n", withdrawAmount, accountBalance)
				fmt.Println("#################################")
			}
		case 4:
			fmt.Println("\n\t#####################################")
			fmt.Println("\tThank you for using Go Bank. Goodbye!")
			fmt.Println("\t#####################################\n")
			return
		default:
			fmt.Println("\n### Invalid choice, please select a valid option. ###")
		}
	}
}