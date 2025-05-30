package main

import "fmt"

type integer int

var x integer
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42

	fmt.Println(x)

	y = int(x) // O tipo customizado e o tipo primitivo são distintos e precisam de conversão explicita

	fmt.Println(y)
}
