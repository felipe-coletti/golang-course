package main

import "fmt"

func x() int {
	return 1
}

func y() (int, string) {
	return 1, "a"
}

func main() {
	fmt.Println(x())
	fmt.Println(y())
}
