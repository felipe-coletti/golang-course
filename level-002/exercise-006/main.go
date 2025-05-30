package main

import (
	"fmt"
)

const (
	_ = 2000 + iota
	OneYearAhead
	TwoYearsAhead
	ThreeYearsAhead
	FourYearsAhead
)

func main() {
	fmt.Println(OneYearAhead, TwoYearsAhead, ThreeYearsAhead, FourYearsAhead)
}
