package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name    string
	Surname string
	Age     int
	Hobbies []string
}

func main() {
	people := []Person{
		{
			Name:    "Ana",
			Surname: "Silva",
			Age:     28,
			Hobbies: []string{"Ler", "Caminhar", "Cozinhar"},
		},
		{
			Name:    "Bruno",
			Surname: "Santos",
			Age:     34,
			Hobbies: []string{"Tocar violão", "Jogar xadrez", "Correr"},
		},
		{
			Name:    "Carla",
			Surname: "Oliveira",
			Age:     31,
			Hobbies: []string{"Viajar", "Fotografar", "Pintar"},
		},
		{
			Name:    "Diego",
			Surname: "Pereira",
			Age:     27,
			Hobbies: []string{"Andar de bicicleta", "Programar", "Assistir filmes"},
		},
		{
			Name:    "Alice",
			Surname: "Martins",
			Age:     30,
			Hobbies: []string{"Desenhar", "Ler fantasia", "Cantar"},
		},
	}

	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Surname < people[j].Surname
		}

		return people[i].Age < people[j].Age
	})

	for _, p := range people {
		sort.Strings(p.Hobbies)
	}

	for _, p := range people {
		fmt.Printf("Nome: %v %v\n", p.Name, p.Surname)
		fmt.Printf("Idade: %dyo\n", p.Age)
		fmt.Println("Hobbies:")

		for i, hobby := range p.Hobbies {
			fmt.Printf("\t%d - %v\n", i+1, hobby)
		}
	}
}
