package main

import "fmt"

// Hello function returns the string "Hello, world!"
func Hello() string {
	return "Hello, world!"
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// HelloYou takes inputString and returns the string "Hello, {inputString}"
func HelloYou(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello())
}
