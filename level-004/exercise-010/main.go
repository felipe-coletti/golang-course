package main

import "fmt"

func main() {
	people := map[string][]string{
		"silva_ana":      {"Ler", "Caminhar", "Cozinhar"},
		"santos_bruno":   {"Tocar violão", "Jogar xadrez", "Correr"},
		"oliveira_carla": {"Viajar", "Fotografar", "Pintar"},
	}

	delete(people, "santos_bruno")

	for key, hobbies := range people {
		fmt.Printf("Pessoa: %v\n", key)

		for i, hobby := range hobbies {
			fmt.Printf("\tHobby %v: %v\n", i, hobby)
		}
	}
}
