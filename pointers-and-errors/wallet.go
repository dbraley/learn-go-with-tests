package wallet

import "fmt"

// Bitcoin type of int
type Bitcoin int

// Wallet manages money
type Wallet struct {
	balance Bitcoin
}

// Deposit adds money to wallet
func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

// Withdraw removes money from wallet
func (w *Wallet) Withdraw(amt Bitcoin) {
	w.balance -= amt
}

// Balance shows the amount of money in the Wallet
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
