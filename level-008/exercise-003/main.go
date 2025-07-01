package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
}

func main() {
	p := Person{
		FirstName: "Ana",
		LastName:  "Silva",
		Age:       28,
		Hobbies:   []string{"Ler", "Caminhar", "Cozinhar"},
	}

	json.NewEncoder(os.Stdout).Encode(p)
}
