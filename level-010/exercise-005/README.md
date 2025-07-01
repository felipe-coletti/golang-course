# Exercise 5

Using this code:

```go
package main

import "fmt"

func main() {
	c := make(chan int)

	v, ok :=
	fmt.Println(v, ok)

	close(c)

	v, ok =
	fmt.Println(v, ok)
}
```

Demonstrate the comma ok idiom.
