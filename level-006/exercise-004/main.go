package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     int
}

func (person Person) nameAndAge() {
	fmt.Printf("%v, %vyo\n", person.name, person.age)
}

func main() {
	person := Person{
		name:    "Ana",
		surname: "Silva",
		age:     28,
	}

	person.nameAndAge()
}
