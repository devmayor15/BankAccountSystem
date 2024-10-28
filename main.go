package main

import "fmt"

//bank account system
type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

//method to deposit money
func (acc *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		acc.Balance += amount
		fmt.Printf("Deposited $%.2f. New Balance: $.%.2f", amount, acc.Balance)
	} else {
		fmt.Println("Deposit amount should be positive")
	}
}

//method to withdraw money
func (acc *BankAccount) Withdraw(amount float64) {
	if amount > acc.Balance {
		fmt.Println("Insufficient funds for withdrawal")
	} else if amount <= 0 {
		fmt.Println("Withdrawal must be positive")
	} else {
		acc.Balance -= amount
		fmt.Printf("Withdrew $%.2f. New Balance: $%.2f", amount, acc.Balance)
	}
}

//method to display the current balance
func (acc *BankAccount) DisplayBalance() {
	fmt.Printf("Current Balance for Account %s: $%.2f\n", acc.AccountNumber, acc.Balance)
}

func main() {
	account := BankAccount{
        AccountNumber: "123456789",
        HolderName:    "Alice Smith",
        Balance:       500.00,
    }
	
	//display balance
	account.DisplayBalance()

	//deposit
	account.Deposit(200.00)

	//withdraw
	account.Withdraw(100.00)

	//attempt to withdraw more than balance
	account.Withdraw(800.00)

	//display balance
	account.DisplayBalance()
}
