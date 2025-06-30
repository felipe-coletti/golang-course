package main

import (
	"fmt"
	"exercise-001/dog"
)

func main() {
	humanYears := 5 // Anos humanos do cachorro
	dogWeight := 10.0 // Peso do cachorro em kg
	dogYears := dog.AgeInDogYears(humanYears, dogWeight) // Converte anos humanos em anos caninos

	fmt.Printf("Um cachorro que possui %d anos humanos possui %d anos caninos.\n", humanYears, dogYears)
}