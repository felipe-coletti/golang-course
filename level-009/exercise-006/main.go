package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS:\t%v\n", runtime.GOOS)
	fmt.Printf("ARCH:\t%v\n", runtime.GOARCH)
}
