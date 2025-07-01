# Exercise 2

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

	bs, err := toJSON(p1)

	fmt.Println(string(bs))
}

func toJSON(a interface{}) []byte {
	bs, _ := json.Marshal(a)
}
```

Create a custom error message using fmt.Errorf().
