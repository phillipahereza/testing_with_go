package banking

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds error
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin of int
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet where money is kept
type Wallet struct {
	balance Bitcoin
}

// Deposit adds money to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += 10
}

// Balance gets wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw removes some money
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
} 
