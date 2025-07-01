# Exercise 1

Make this code work using a self-executing anonymous function and buffer:

```go
package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
```
