package main

import "fmt"

type Person struct {
	name            string
	surname         string
	favoriteFlavors []string
}

func main() {
	people := map[string]Person{
		"Silva": {
			name:            "Ana",
			surname:         "Silva",
			favoriteFlavors: []string{"Morango", "Chocolate"},
		},
		"Santos": {
			name:            "Bruno",
			surname:         "Santos",
			favoriteFlavors: []string{"Baunilha", "Menta"},
		},
	}

	for _, person := range people {
		fmt.Printf("%v %v\n", person.name, person.surname)
		fmt.Println("Sabores favoritos:")

		for i, flavor := range person.favoriteFlavors {
			fmt.Printf("\t%d - %v\n", i+1, flavor)
		}
	}
}
