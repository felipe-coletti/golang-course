package main

import "fmt"

func main() {
	count := 0

	increment := func() int {
		count++

		return count
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
