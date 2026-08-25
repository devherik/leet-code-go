package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrInsufficientFounds = errors.New("insufficient founds")

// account
type Account struct {
	mu      sync.RWMutex
	balance float64
}

func NewAcount(initialBalance float64) *Account {
	return &Account{
		balance: initialBalance,
	}
}

// deposit
func (a *Account) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock() // gararntee the release even on panic

	a.balance += amount
}

// withdraw
func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance < amount {
		return ErrInsufficientFounds
	}
	a.balance -= amount
	return nil
}

// balance guarantee a correct read of the balance by locking the read
func (a *Account) Balance() float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.balance
}

func mainMutex() {
	acc := NewAcount(100.0)
	var wg sync.WaitGroup

	for range 50 {
		// add two goroutines to the waiting group in every loop
		wg.Add(2)
		go func() {
			defer wg.Done() // mark as done when the function exits
			acc.Deposit(10.0)
		}()
		go func() {
			defer wg.Done()
			_ = acc.Withdraw(5.0)
		}()
	}

	// wait for all goroutines to finish
	wg.Wait()
	fmt.Printf("Saldo final: R$ %.2f\n", acc.Balance())
}
