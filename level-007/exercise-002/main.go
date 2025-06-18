package main

import "fmt"

type Person struct {
	name    string
	surname string
}

func (person *Person) setName(newName string) {
	(*person).name = newName
}

func main() {
	person := Person{
		name:    "Ana",
		surname: "Silva",
	}

	(&person).setName("Bruno")

	fmt.Println(person.name)
}
