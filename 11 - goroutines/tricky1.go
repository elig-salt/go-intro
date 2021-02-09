package main

import "fmt"

func PrintHi() {
	fmt.Println("Hi")
}

func Tricky1() {
	go PrintHi()
	fmt.Println("Hi2")
}
