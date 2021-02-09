package main

import "fmt"

func main() {
	// No name for this struct
	// But we need to declare the type here
	bio := struct {
		firstName string
		friends   map[string]int
		favDrinks []string
	}{
		firstName: "Steven",
		friends: map[string]int{
			"Tim":   12345678,
			"Abdul": 34789657,
			"Kally": 28993332,
		},
		favDrinks: []string{
			"Pepsi",
			"Cocacola",
		},
	}
	fmt.Println(bio.firstName)

	for k, v := range bio.friends {
		fmt.Println(k, v)
	}

	for k, v := range bio.favDrinks {
		fmt.Println(k, v)
	}
}
