package wallet

import "fmt"

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

//Withdraw function allows withdrawal of a specific amount of BTC
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
