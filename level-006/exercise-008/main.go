package main

import "fmt"

var sayHello = func() func() {
	return func() {
		fmt.Println("Hello, world!")
	}
}()

func main() {
	sayHello()
}
