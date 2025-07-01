# Exercise 1

Using the code below, use marshal to transform `[]user` into JSON.

```go
package main

import "fmt"

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	u2 := user{
		first: "Moneypenny",
		age:   27,
	}

	u3 := user{
		first: "M",
		age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
}
```
