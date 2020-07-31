package main

import "fmt"

// Hello function returns the string "Hello, world!"
func Hello() string {
	return "Hello, world!"
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// HelloYou takes inputString and returns the string "Hello, {inputString}"
func HelloYou(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := ""

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name

}

func main() {
	fmt.Println(Hello())
}
