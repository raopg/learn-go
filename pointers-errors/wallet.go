package wallet

import (
	"errors"
	"fmt"
)

//Bitcoin type to represent the contents of the balance
type Bitcoin int

//Wallet struct holds all data pertaining to the BTC wallet
type Wallet struct {
	balance Bitcoin
}

//Deposit function takes an amout and adds it to the balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += 10
}

//Balance function returns the balance of the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Stringer struct to extend the string representation of Bitcoin struct
type Stringer interface {
	String() string
}

//String function for a BTC string for the wallet
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//ErrInsufficientFunds holds the exact error message for when overdrafting withdrawals
var ErrInsufficientFunds = errors.New("Cannot withdraw; Insufficient funds")

//Withdraw function allows withdrawal of a specific amount of BTC
func (w *Wallet) Withdraw(amount Bitcoin) error {

	//We want to throw an error if the amount is greater than current balance
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
