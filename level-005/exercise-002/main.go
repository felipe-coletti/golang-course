package main

import "fmt"

type person struct {
	name            string
	surname         string
	favoriteFlavors []string
}

func main() {
	people := map[string]person{
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

	for _, p := range people {
		fmt.Printf("%v %v\n", p.name, p.surname)
		fmt.Println("Sabores favoritos:")

		for i, flavor := range p.favoriteFlavors {
			fmt.Printf("\t%d - %v\n", i+1, flavor)
		}
	}
}
