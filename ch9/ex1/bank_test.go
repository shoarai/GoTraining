// Copyright © 2016 shoarai

package bank_test

import (
	"fmt"
	"testing"

	"github.com/shoarai/GoTraining/ch9/ex1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Withdraw(400)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	if got := bank.Balance(); got < 0 {
		t.Errorf("Balance = %d, want more than zero", got)
	}
}
