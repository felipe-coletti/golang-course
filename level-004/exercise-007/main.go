package main

import "fmt"

func main() {
	people := [][]string{
		{"Ana", "Silva", "Ler"},
		{"Bruno", "Santos", "Tocar violão"},
		{"Carla", "Oliveira", "Viajar"},
	}

	for i, person := range people {
		fmt.Printf("Pessoa %d:\n", i+1)
		fmt.Printf("\tNome: %s\n", person[0])
		fmt.Printf("\tSobrenome: %s\n", person[1])
		fmt.Printf("\tHobby favorito: %s\n", person[2])
	}
}
