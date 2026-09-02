package main

type Patient struct {
	name               string
	registrationDone   bool
	doctorCheckoutDone bool
	medicineDone       bool
	paymentDone        bool
}

type Departement interface {
	execute(*Patient)
	setNext(Departement)
}
