package main

import "fmt"

type integer int

var x integer
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42

	fmt.Println(x)

	y = int(x) // Custom type and primitive type are distinct and need explicit conversion

	fmt.Println(y)
}
