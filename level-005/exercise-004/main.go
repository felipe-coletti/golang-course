package main

import "fmt"

func main() {
	people := []struct {
		firstName string
		lastName  string
	}{
		{
			firstName: "Ana",
			lastName:  "Silva",
		},
		{
			firstName: "Bruno",
			lastName:  "Santos",
		},
	}

	for _, person := range people {
		fmt.Printf("%v %v\n", person.firstName, person.lastName)
	}
}
