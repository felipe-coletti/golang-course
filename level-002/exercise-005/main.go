package main

import (
	"fmt"
)

func main() {
	rawString := `Essa é uma raw string.
Ela pode conter ', "
e múltiplas linhas,
mas não permite interpolação.`

	fmt.Println(rawString)
}
