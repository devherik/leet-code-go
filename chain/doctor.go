package main

import "fmt"

type Doctor struct {
	next Departement
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckoutDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckoutDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Departement) {
	d.next = next
}
