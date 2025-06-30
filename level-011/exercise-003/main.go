package main

import "fmt"

type customError struct {
	msg string
}

func (e customError) Error() string {
	return e.msg
}

func printError(e error) {
	fmt.Println(e) // Equivalente a fmt.Println(e.Error())
}

func main() {
	err := customError{"¯\\_(ツ)_/¯"}

	printError(err)

}
