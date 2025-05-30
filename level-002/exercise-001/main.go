package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := int64(100)

	dec := strconv.FormatInt(num, 10)
	bin := strconv.FormatInt(num, 2)
	hex := strconv.FormatInt(num, 16)

	fmt.Println("Decimal: ", dec)
	fmt.Println("Binário: ", bin)
	fmt.Println("Hexadecimal: ", hex)
}
