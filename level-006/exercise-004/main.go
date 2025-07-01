package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) nameAndAge() {
	fmt.Printf("%v %v, %dyo\n", p.firsName, p.lastName, p.age)
}

func main() {
	p := person{
		firsName: "Ana",
		lastName: "Silva",
		age:      28,
	}

	p.nameAndAge()
}
