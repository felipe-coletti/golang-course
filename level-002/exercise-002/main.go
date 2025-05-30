package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	equal := a == b
	different := a != b
	greaterThan := a > b
	greaterOrEqual := a >= b
	lessThan := a < b
	lessOrEqual := a <= b

	fmt.Println("a é igual a b: ", equal)
	fmt.Println("a é diferente de b: ", different)
	fmt.Println("a é maior que b: ", greaterThan)
	fmt.Println("a é maior ou igual a b: ", greaterOrEqual)
	fmt.Println("a é menor que b: ", lessThan)
	fmt.Println("a é menor ou igual a b: ", lessOrEqual)
}
