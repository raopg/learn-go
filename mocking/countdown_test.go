package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("basic countdown test", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer)

		output := buffer.String()
		expected := `3
2
1
Go!`

		assertStringEquals(t, output, expected)

	})
}

func assertStringEquals(t *testing.T, output, expected string) {
	if output != expected {
		t.Errorf("Expected: %q\nActual Output: %q", expected, output)
	}
}
