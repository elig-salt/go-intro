package main

import "fmt"

func Tricky2() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
