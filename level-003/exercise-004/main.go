package main

import (
	"fmt"
	"time"
)

func main() {
	var birthYear int
	currentYear := time.Now().Year()

	fmt.Print("Digite seu ano de nascimento: ")
	fmt.Scanln(&birthYear)

	year := birthYear

	for {
		fmt.Println(year)

		year++

		if year > currentYear {
			break
		}
	}
}
