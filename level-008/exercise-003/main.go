package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name    string
	Surname string
	Age     int
	Hobbies []string
}

func main() {
	p := Person{
		Name:    "Ana",
		Surname: "Silva",
		Age:     28,
		Hobbies: []string{"Ler", "Caminhar", "Cozinhar"},
	}

	json.NewEncoder(os.Stdout).Encode(p)
}
