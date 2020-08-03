package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test greet basic", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Greet(&buffer, "Prateek")

		output := buffer.String()
		expected := "Hello, Prateek"

		assertStringEquals(t, output, expected)

	})
}

func assertStringEquals(t *testing.T, output, expected string) {
	t.Helper()

	if output != expected {
		t.Errorf("Expected: %q\nActual Output: %q", expected, output)
	}
}
