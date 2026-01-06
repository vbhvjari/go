package main

import "fmt"

type UPI struct {
	upiid  string
	amount int
}
type Card struct {
	transid string
	amount  int
}

func (u UPI) pay() {
	fmt.Printf("\nPayment via UPI with upiid : %s and amount: %d", u.upiid, u.amount)
}

func (c Card) pay() {
	fmt.Printf("\nPayment via Card with transid : %s and amount: %d", c.transid, c.amount)
}

type Payment interface {
	pay()
}

func processPayment(p Payment) {
	p.pay()
}

func main() {
	u1 := UPI{upiid: "abc1234", amount: 100}
	c1 := Card{transid: "xyz345", amount: 200}
	processPayment(u1)
	processPayment(c1)
}
