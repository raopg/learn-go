package main

import "fmt"

// Hello function returns the string "Hello, world!"
func Hello() string {
	return "Hello, world!"
}

const englishHelloPrefix = "Hello, "

// HelloYou takes inputString and returns the string "Hello, {inputString}"
func HelloYou(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello())
}
