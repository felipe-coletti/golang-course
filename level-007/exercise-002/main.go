package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p *person) setName(newName string) {
	(*p).firstName = newName
}

func main() {
	p := person{
		firstName: "Ana",
		lastName:  "Silva",
	}

	(&p).setName("Bruno")

	fmt.Println(p.firstName)
}
