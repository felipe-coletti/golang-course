package main

import "fmt"

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func receive(channel <-chan int) {
	for v := range channel {
		fmt.Println(v)
	}
}

func main() {
	c := gen()

	receive(c)
}
