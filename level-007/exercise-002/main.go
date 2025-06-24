package main

import "fmt"

type person struct {
	name    string
	surname string
}

func (p *person) setName(newName string) {
	(*p).name = newName
}

func main() {
	p := person{
		name:    "Ana",
		surname: "Silva",
	}

	(&p).setName("Bruno")

	fmt.Println(p.name)
}
