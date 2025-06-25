package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	words := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(nums)
	fmt.Println(words)

	sort.Ints(nums)
	sort.Strings(words)

	fmt.Println(nums)
	fmt.Println(words)
}
