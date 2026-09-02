package main

import "fmt"

type Cashier struct {
	next Departement
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) setNext(next Departement) {
	c.next = next
}
