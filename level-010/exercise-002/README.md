# Exercise 2

Make this code work:

```go
package main

import "fmt"

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)
	fmt.Printf("cs\t%T\n", cs)
}
```
