package main

import "fmt"

func main() {
	people := map[string][]string{
		"silva_ana":      {"Ler", "Caminhar", "Cozinhar"},
		"santos_bruno":   {"Tocar violão", "Jogar xadrez", "Correr"},
		"oliveira_carla": {"Viajar", "Fotografar", "Pintar"},
	}

	people["pereira_diego"] = []string{"Andar de bicicleta", "Programar", "Assistir filmes"}

	for key, hobbies := range people {
		fmt.Printf("Pessoa: %s\n", key)

		for i, hobby := range hobbies {
			fmt.Printf("\tHobby %d: %s\n", i, hobby)
		}
	}
}
