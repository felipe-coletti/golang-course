# Exercise 1

Using this code:

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, _ := json.Marshal(p1)

	fmt.Println(string(bs))
}
```

Remove the underscore and check and handle the error appropriately.
