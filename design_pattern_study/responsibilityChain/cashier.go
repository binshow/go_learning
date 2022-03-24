package main

import "fmt"

// 收银

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
	c.next.execute(p)
}

func (c *cashier) setNext(next department) {
	c.next = next
}