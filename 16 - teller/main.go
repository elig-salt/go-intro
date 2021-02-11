package main

import (
	"fmt"
	"sync"
	"time"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			Deposit(1)
			time.Sleep(2 * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		for range ticker.C {
			balance := Balance()
			fmt.Println(balance)
		}
	}()

	wg.Wait()
}
