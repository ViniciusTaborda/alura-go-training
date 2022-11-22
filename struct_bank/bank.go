package main

import (
	"errors"
	"fmt"
	"math"
)

var negativeAmount error = errors.New("Amount can not be negative.")

type BankAccount interface {
	withdraw(amount float64) (float64, error)
}

type AccountHolder struct {
	name string
	age  int
}

type CheckingAccount struct {
	accountHolder AccountHolder
	agencyNumber  int
	accountNumber int
	balance       float64
}
type SavingsAccount struct {
	accountHolder AccountHolder
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {

	firstAccount := CheckingAccount{
		accountHolder: AccountHolder{
			name: "Vinicius",
			age:  21,
		},
		agencyNumber:  1365,
		accountNumber: 295,
		balance:       249.34,
	}

	fmt.Printf("firstAccount: %+v\n", firstAccount)

	currBalance, err := firstAccount.withdraw(-500.0)

	fmt.Println(payDebt(&firstAccount, 150))

	fmt.Println(currBalance, err)

}

func payDebt(account BankAccount, value float64) (err error) {
	_, err = account.withdraw(value)
	return
}

func (checkingAcc *CheckingAccount) withdraw(amount float64) (float64, error) {

	var err error

	if !(math.Signbit(amount)) {
		if amount < checkingAcc.balance {
			checkingAcc.balance -= amount
		} else {
			err = errors.New("Requested withdraw amount is grater than balance.")
		}
	} else {
		err = negativeAmount
	}

	return checkingAcc.balance, err
}

func (savingsAcc *SavingsAccount) withdraw(amount float64) (float64, error) {

	var err error

	if !(math.Signbit(amount)) {
		if amount < savingsAcc.balance {
			savingsAcc.balance -= amount
		} else {
			err = errors.New("Requested withdraw amount is grater than balance.")
		}
	} else {
		err = negativeAmount
	}

	return savingsAcc.balance, err
}

func (checkingAcc *CheckingAccount) deposit(amount float64) (float64, error) {

	var err error

	if !(math.Signbit(amount)) {
		checkingAcc.balance += amount
	} else {
		err = negativeAmount
	}

	return checkingAcc.balance, err
}
func (savingsAcc *SavingsAccount) deposit(amount float64) (float64, error) {

	var err error

	if !(math.Signbit(amount)) {
		savingsAcc.balance += amount
	} else {
		err = negativeAmount
	}

	return savingsAcc.balance, err
}

func (checkingAcc *CheckingAccount) getBalance() float64 {
	return checkingAcc.balance
}
