package main

import "fmt"

func main() {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	runesSlice := []rune(chars)

	for _, currentRune := range runesSlice {
		fmt.Println(currentRune)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", currentRune)
		}
	}
}
