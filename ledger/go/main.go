package main

import (
	"fmt"
	"time"
)

type Posting struct {
	AccountID string
	Amount    int64 // Storing money in cents/integer to avoid floating-point errors
}

type Transaction struct {
	ID        string
	Debits    []Posting
	Credits   []Posting
	Timestamp time.Time
}

type Ledger struct {
	transactions []Transaction
}

func (l *Ledger) PostTransaction(tx Transaction) error {
	var totalDebits int64
	var totalCredits int64

	for _, d := range tx.Debits {
		totalDebits += d.Amount
	}

	for _, c := range tx.Credits {
		totalCredits += c.Amount
	}

	if totalDebits != totalCredits {
		return fmt.Errorf("Transaction %s is unbalanced", tx.ID)
	}

	l.transactions = append(l.transactions, tx)
	return nil
}

func (l *Ledger) GetBalance(accountID string) int64 {
	var balance int64

	for _, tx := range l.transactions {
		for _, d := range tx.Debits {
			if d.AccountID == accountID {
				balance -= d.Amount
			}
		}

		for _, c := range tx.Credits {
			if c.AccountID == accountID {
				balance += c.Amount
			}
		}
	}

	return balance
}

func main() {
	ledger := &Ledger{}

	// Move $10.50 (1050 cents) from Revenue to Cash Reserve
	tx := Transaction{
		ID:        "tx_999",
		Timestamp: time.Now(),
		Debits: []Posting{
			{AccountID: "revenue_acc", Amount: 1050},
		},
		Credits: []Posting{
			{AccountID: "cash_reserve", Amount: 1050},
		},
	}

	err := ledger.PostTransaction(tx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Revenue Acc Balance: %d cents\n", ledger.GetBalance("revenue_acc"))   // -1050
	fmt.Printf("Cash Reserve Balance: %d cents\n", ledger.GetBalance("cash_reserve")) // 1050
}
