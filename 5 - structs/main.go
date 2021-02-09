package main

import "fmt"

type person struct {
	firstName string `json:"firstName" yaml:"firstName"` // Tagging can be added for serialization
	lastName  string
	age       int
}

func main() {
	//Assigning values to the fields in the person struct:
	p1 := person{
		firstName: "Carl",
		lastName:  "Eschenbach",
		age:       54,
	}

	fmt.Println("He is the man: ", p1)

	type animal struct {
		name            string
		characteristics []string
	}

	animal1 := animal{
		name: "Lion",
		characteristics: []string{"Eats human",
			"Wild Animal",
			"King of the jungle",
		},
	}

	//We use dot(.) to acces each field in the struct
	fmt.Println("Animal name: ", animal1.name)
	for _, v := range animal1.characteristics {
		fmt.Printf("\t %v\n", v)
	}

	// Nested Structs

	//A herbivore is an animal, so it can have the animal struct as a field
	type herbivore struct {
		animal
		eatHuman bool
	}

	herb := herbivore{
		animal: animal{
			name: "Goat",
			characteristics: []string{"Lacks sense",
				"Lazy",
				"Eats grass",
			},
		},
		eatHuman: false, //maybe
	}

	fmt.Println("\nThis animal:")
	//We use dot(.) to acces each field in the struct
	fmt.Println("Eats human? ", herb.eatHuman)
	// Promotion
	fmt.Println("Animal name:", herb.name) // <<====== No need for Animal!!
	fmt.Println("Characteristics:")
	for _, v := range herb.characteristics {
		fmt.Printf("\t %v\n", v)
	}
}
