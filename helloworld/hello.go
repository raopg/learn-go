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
	return getGreetingPrefix(language) + name

}

//getGreetingPrefix determines the greeting prefix based on the input language
func getGreetingPrefix(language string) (prefix string) {
	// Notice how we a named return value. Here we dont have to return the value specifically
	// The prefix var is also initialized to a default val. For strings, this is ""
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello())
}
