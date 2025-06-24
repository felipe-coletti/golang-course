package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (p person) nameAndAge() {
	fmt.Printf("%v, %vyo\n", p.name, p.age)
}

func main() {
	p := person{
		name:    "Ana",
		surname: "Silva",
		age:     28,
	}

	p.nameAndAge()
}
