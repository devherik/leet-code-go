package main

import (
	"testing"
	"time"
)

func TestPostTransaction(t *testing.T) {
	tests := []struct {
		name        string
		tx          Transaction
		expectError bool
	}{
		{
			name: "Valid double-entry transaction",
			tx: Transaction{
				ID:        "tx_ok",
				Timestamp: time.Now(),
				Debits: []Posting{
					{AccountID: "1", Amount: 100},
				},
				Credits: []Posting{
					{AccountID: "2", Amount: 100},
				},
			},
			expectError: false,
		},
		{
			name: "Invalid transaction",
			tx: Transaction{
				ID:        "tx_bad",
				Timestamp: time.Now(),
				Debits: []Posting{
					{AccountID: "1", Amount: 100},
				},
				Credits: []Posting{
					{AccountID: "2", Amount: 99},
				},
			},
			expectError: true,
		},
	}

	for _, tx := range tests {
		t.Run(tx.name, func(t *testing.T) {
			legder := &Ledger{}
			err := legder.PostTransaction(tx.tx)

			if (err != nil) != tx.expectError {
				t.Errorf("Transaction %s: got error %v, want error %v", tx.name, err != nil, tx.expectError)
			}
		})
	}
}

func TestGetBalanceAggregation(t *testing.T) {
	ledger := &Ledger{}

	tx1 := Transaction{
		ID: "tx_01",
		Debits: []Posting{
			{AccountID: "main", Amount: 100},
		},
		Credits: []Posting{
			{AccountID: "secondary", Amount: 100},
		},
		Timestamp: time.Now(),
	}

	tx2 := Transaction{
		ID: "tx_02",
		Debits: []Posting{
			{AccountID: "secondary", Amount: 200},
		},
		Credits: []Posting{
			{AccountID: "main", Amount: 200},
		},
		Timestamp: time.Now(),
	}

	_ = ledger.PostTransaction(tx1)
	_ = ledger.PostTransaction(tx2)

	if got := ledger.GetBalance("main"); got != 100 {
		t.Errorf("Main account balance: got %d, want 100", got)
	}

	if got := ledger.GetBalance("secondary"); got != -100 {
		t.Errorf("Secondary account balance: got %d, want -100", got)
	}
}
