package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for _, item := range array {
		fmt.Println(item)
	}

	fmt.Printf("%T", array)
}
