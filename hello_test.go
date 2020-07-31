package main

import "testing"

func TestHello(t *testing.T) {
	output := Hello()
	expected := "Hello, world!"

	if output != expected {
		t.Errorf("Output: %q \nExpected Output: %q", output, expected)
	}
}

func TestHelloYou(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, output, expected string) {
		t.Helper()
		if output != expected {
			t.Errorf("Output: %q \nExpected Output: %q", output, expected)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		output := HelloYou("Prateek", "")
		expected := "Hello, Prateek"
		assertCorrectMessage(t, output, expected)
	})

	t.Run("Saying Hello world when no person is given", func(t *testing.T) {
		output := HelloYou("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, output, expected)
	})

	t.Run("Saying hello in Spanish", func(t *testing.T) {
		output := HelloYou("Jeff", "Spanish")
		expected := "Hola, Jeff"
		assertCorrectMessage(t, output, expected)
	})

	t.Run("Saying hello in French", func(t *testing.T) {
		output := HelloYou("Oui Oui", "French")
		expected := "Bonjour, Oui Oui"
		assertCorrectMessage(t, output, expected)
	})
}
