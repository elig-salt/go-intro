package main

import (
	"fmt"
	"sync"
)

func main() {
	// Simple()
	// Blocking()
	// Buffered()
	// Nil()
	// BrokenFor()
	// Range()
	// Closing()
	Broadcast()
}

func Simple() {
	someChan := make(chan int)
	go func() { someChan <- 1 }()
	a := <-someChan
	fmt.Println(a)
}

func Blocking() {
	someChan := make(chan bool)

	// 1. Will Block
	// someChan <- true

	// 2. Wil block now too
	// <-someChan

	go func() {
		someChan <- true
	}()
	fmt.Println(<-someChan)
}

func Buffered() {
	someChan := make(chan bool, 1)
	someChan <- true // No block, we have 1 slot
	fmt.Println(<-someChan)

	someChan = make(chan bool, 2)
	someChan <- true // No block, we have 1 slot
	someChan <- true // No block, we have 1 slot
	fmt.Println(<-someChan)
}

func Nil() {
	var nilChan chan bool
	<-nilChan // Nil channels will block forever
}

func BrokenFor() {
	someChan := make(chan int)
	go func() {
		// Infinite loop !
		for {
			fmt.Println(<-someChan)
		}
	}()
	someChan <- 1
	someChan <- 2
	someChan <- 3
	// No print for 3?
}

func Range() {
	someChan := make(chan int)
	go func() {
		someChan <- 1
		someChan <- 2
		someChan <- 3
		close(someChan)
	}()
	for val := range someChan {
		fmt.Println(val)
	}
}

func Closing() {
	// Closing a channel
	// 1. Writing to a closed channel will panic
	someChan := make(chan int)
	close(someChan)
	// someChan <- 3 // panic

	// 2. Reading from a closed channel will always succeed (unless in range)
	someChan = make(chan int)
	close(someChan)
	// First will be zero-value
	// Second is always a boolean indicating if the channel is open
	val, ok := <-someChan
	fmt.Printf("Val is:  %v, is open: %v\n", val, ok)

	// RULE OF THUMB:
	// ONLY PRODUCER CLOSES THE CHANNEL
}

func Broadcast() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go func(c <-chan int) {
		for i := range c {
			fmt.Println("1st goroutine:", i)
		}
		// Next line happens only if c is closed
		wg.Done()
	}(ch)

	go func(c <-chan int) {
		for i := range c {
			fmt.Println("2st goroutine:", i)
		}
		// Next line happens only if c is closed
		wg.Done()
	}(ch)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // This line caused both goroutines exit the range
	wg.Wait()
}
