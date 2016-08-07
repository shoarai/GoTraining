// Copyright © 2016 shoarai

package bank_test

import (
	"fmt"
	"testing"

	"github.com/shoarai/GoTraining/ch9/ex1"
)

func TestBank(t *testing.T) {
	tests := []struct {
		inputs []int
		want   int
	}{
		// Balance is stocked
		{[]int{100, 200}, 300},
		{[]int{0, 600, 100}, 1000},
	}

	for _, test := range tests {
		done := make(chan struct{})
		for i := 0; i < len(test.inputs); i++ {
			go func(input int) {
				bank.Deposit(input)
				fmt.Println("=", bank.Balance())
				done <- struct{}{}
			}(test.inputs[i])
		}

		// Wait for both transactions.
		for i := 0; i < len(test.inputs); i++ {
			<-done
		}

		if got := bank.Balance(); got != test.want {
			t.Errorf("Balance = %d, want %d", got, test.want)
		}
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		input   int
		want    bool
		balance int
	}{
		// Balance is stocked
		{200, true, 800},
		{1000, false, 800},
		{800, true, 0},
	}

	for _, test := range tests {
		isSuccessWithdraw := make(chan bool)
		go func(input int) {
			g := bank.Withdraw(input)
			fmt.Println("=", bank.Balance())
			isSuccessWithdraw <- g
		}(test.input)

		if got := <-isSuccessWithdraw; got != test.want {
			t.Errorf("Withdraw(%d) = %t", test.input, got)
		}

		if balance := bank.Balance(); balance != test.balance {
			t.Errorf("Balance = %d, want %d", balance, test.balance)
		}
	}
}
