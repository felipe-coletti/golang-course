package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (p person) nameAndAge() {
	fmt.Printf("%v %v, %dyo\n", p.name, p.surname, p.age)
}

func main() {
	p := person{
		name:    "Ana",
		surname: "Silva",
		age:     28,
	}

	p.nameAndAge()
}
