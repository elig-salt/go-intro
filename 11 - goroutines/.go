package main

import (
	"fmt"
	"sync"
)

func Tricky3() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Print(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

// Suggest a change !
