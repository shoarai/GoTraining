// Copyright © 2016 shoarai

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)  // send amount to deposit
var balances = make(chan int)  // receive balance
var withdraws = make(chan int) // send amount to withdraw

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	balance := Balance()
	if balance < amount {
		return false
	}
	withdraws <- amount
	return true
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			balance -= amount
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
