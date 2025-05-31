package main

import (
	"fmt"
	"strings"
)

func main() {
	var favoriteSport string

	fmt.Print("Digite seu esporte favorito: ")
	fmt.Scanln(&favoriteSport)

	favoriteSport = strings.ToLower(favoriteSport)

	switch favoriteSport {
	case "futebol":
		fmt.Println("Futebol é um esporte muito interessante.")
	case "basquete":
		fmt.Println("Basquete parece ser bem divertido.")
	case "críquete":
		fmt.Println("Críquete é um esporte bem diferente.")
	case "rúgbi":
		fmt.Println("Rúgbi é cheio de ação e energia.")
	case "tênis":
		fmt.Println("Tênis exige muita habilidade e concentração.")
	default:
		fmt.Println("Não conheço esse esporte.")
	}
}
