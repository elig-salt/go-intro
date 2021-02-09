package main

import "fmt"

var sampleMap map[string]int

func main() {
	sampleMap = map[string]int{
		"James": 50,
		"Ali":   39,
	}

	currency := map[string]string{
		"AUD": "Australia Dollar",
		"GBP": "Great Britain Pound",
		"JPY": "Japan Yen",
		"CHF": "Switzerland Franc",
	}

	//a. Adding to the map:
	currency["USD"] = "USA Dollar"

	fmt.Println("Currency with USD added: ", currency)

	//b. Remove from the map:
	delete(currency, "GBP")
	fmt.Println("Currency with GBP deleted: ", currency)

	//c. Replacing one entry with another:
	currency["AUD"] = "New Zealand Dollar"
	fmt.Println("Currency with AUD value replaced with NZD: ", currency)

	//Ranging through the map:
	for key, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", key, value)
	}
}
