package main

import "fmt"

func main() {
	people := []struct {
		name    string
		surname string
	}{
		{
			name:    "Ana",
			surname: "Silva",
		},
		{
			name:    "Bruno",
			surname: "Santos",
		},
	}

	for _, person := range people {
		fmt.Printf("%v %v\n", person.name, person.surname)
	}
}
