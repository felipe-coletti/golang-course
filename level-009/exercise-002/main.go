package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p *person) say(phrase string) {
	fmt.Printf("%v %v: %v\n", (*p).firstName, (*p).lastName, phrase)
}

type human interface {
	say(string)
}

func say(h human, phrase string) {
	h.say(phrase)
}

func main() {
	p := person{
		firstName: "Ana",
		lastName:  "Silva",
	}

	p.say("Olá!")    // Funciona porque o compilador automaticamente faz (&p).say()
	(&p).say("Olá!") // Funciona diretamente
	say(&p, "Olá!")  // Funciona porque &p implementa `human`
}
