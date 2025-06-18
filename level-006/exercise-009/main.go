package main

import (
	"fmt"
	"strings"
)

func formatMessage(msg string, formatter func(string) string) string {
	return formatter(msg)
}

func toUpperCase(str string) string {
	return strings.ToUpper(str)
}

func toLowerCase(str string) string {
	return strings.ToLower(str)
}

func main() {
	message := "Hello, World!"

	fmt.Println(formatMessage(message, toUpperCase))
	fmt.Println(formatMessage(message, toLowerCase))
}
