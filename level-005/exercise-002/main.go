package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favoriteFlavors []string
}

func main() {
	people := map[string]person{
		"Silva": {
			firstName:       "Ana",
			lastName:        "Silva",
			favoriteFlavors: []string{"Morango", "Chocolate"},
		},
		"Santos": {
			firstName:       "Bruno",
			lastName:        "Santos",
			favoriteFlavors: []string{"Baunilha", "Menta"},
		},
	}

	for _, p := range people {
		fmt.Printf("%v %v\n", p.firstName, p.lastName)
		fmt.Println("Sabores favoritos:")

		for i, flavor := range p.favoriteFlavors {
			fmt.Printf("\t%d - %v\n", i+1, flavor)
		}
	}
}
