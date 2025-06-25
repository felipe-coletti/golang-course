package main

import "fmt"

type person struct {
	name    string
	surname string
}

func (p *person) say(phrase string) {
	fmt.Printf("%v %v: %v\n", (*p).name, (*p).surname, phrase)
}

type human interface {
	say(string)
}

func say(h human, phrase string) {
	h.say(phrase)
}

func main() {
	p := person{
		name:    "Ana",
		surname: "Silva",
	}

	p.say("Olá!")    // Funciona porque o compilador automaticamente faz (&p).say()
	(&p).say("Olá!") // Funciona diretamente
	say(&p, "Olá!")  // Funciona porque &p implementa `human`
}
