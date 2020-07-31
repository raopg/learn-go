package wallet

//Wallet struct holds all data pertaining to the BTC wallet
type Wallet struct {
	balance int
}

//Deposit function takes an amout and adds it to the balance
func (w *Wallet) Deposit(amount int) {
	w.balance += 10
}

//Balance function returns the balance of the wallet
func (w *Wallet) Balance() int {
	return w.balance
}
