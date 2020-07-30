package main

import "fmt"

// Hello function returns the string "Hello, world!"
func Hello() string {
	return "Hello, world!"
}

// HelloYou takes inputString and returns the string "Hello, {inputString}"
func HelloYou(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello())
}
