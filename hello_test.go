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
	output := HelloYou("Prateek")
	expected := "Hello, Prateek"

	if output != expected {
		t.Errorf("Output: %q \nExpected Output: %q", output, expected)
	}
}
