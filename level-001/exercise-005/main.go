package main

import "fmt"

type integer int

var x integer

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42

	fmt.Println(x)
}
