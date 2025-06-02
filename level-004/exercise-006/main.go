package main

import "fmt"

func main() {
	states := make([]string, 0, 26)

	states = append(states, "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins")

	fmt.Println(len(states), cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}
}
