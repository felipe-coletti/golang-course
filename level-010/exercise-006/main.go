package main

import "fmt"

func fill(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}

func main() {
	c := make(chan int)

	go fill(c)

	for v := range c {
		fmt.Println(v)
	}
}
