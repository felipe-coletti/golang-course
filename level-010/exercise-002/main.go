package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v := <-c

	fmt.Println(v)

	fmt.Printf("%v\t%T\n", c, c)
}
