package main

import (
	"fmt"
)

func main() {
	Simple()
}

func Simple() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	exitCh := make(chan interface{})

	go func() {
		ch1 <- 1
		ch2 <- 2
		exitCh <- 1
	}()

	for {
		select {
		case i := <-ch1:
			fmt.Println("Ch2: ", i)
		case i := <-ch2:
			fmt.Println("Ch1: ", i)
		case <-exitCh:
			return
		}
	}
}
