package main

import (
	"fmt"
)

func main() {
	waterLitersDrunk := 1.5

	if waterLitersDrunk < 2 {
		fmt.Println("Beba mais água.")
	} else if waterLitersDrunk == 2 {
		fmt.Println("Parabens! Você bebeu a quantidade adequada de água (na média) para um dia.")
	} else {
		fmt.Println("Calma lá. É bom beber bastante água, mas o excesso também pode te fazer mal.")
	}
}
