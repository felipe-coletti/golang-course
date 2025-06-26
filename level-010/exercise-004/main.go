package main

import "fmt"

func gen(c1 chan<- int) <-chan int {
	c2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c2 <- i
		}

		close(c2)

		c1 <- 0
	}()

	return c2
}

func receive(c1 chan int, c2 <-chan int) {
	for {
		select {
		case v := <-c2:
			fmt.Println(v)
		case <-c1:
			return
		}
	}
}

func main() {
	c1 := make(chan int)
	c2 := gen(c1)

	receive(c1, c2)
}
