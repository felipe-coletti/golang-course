# Exercise 3

Using this code:

```go
package main

import "fmt"

func gen() <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}

func main() {
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}
```

Use a for range loop to demonstrate the channel values.
