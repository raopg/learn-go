package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("basic countdown test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		output := buffer.String()
		expected := `3
2
1
Go!`

		assertStringEquals(t, output, expected)

		assertSleeperCalls(t, spySleeper.Calls, 4)

	})
}

func assertStringEquals(t *testing.T, output, expected string) {
	if output != expected {
		t.Errorf("Expected: %q\nActual Output: %q", expected, output)
	}
}

func assertSleeperCalls(t *testing.T, actualCalls, expectedCalls int) {
	if actualCalls != expectedCalls {
		t.Errorf("Expected %d sleep calls\nGot %d sleep calls", expectedCalls, actualCalls)
	}
}
