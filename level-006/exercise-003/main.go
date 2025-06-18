package main

import "fmt"

func main() {
	defer fmt.Println("Primeiro print")
	defer fmt.Println("Segundo print")
	fmt.Println("Terceiro print")
}
