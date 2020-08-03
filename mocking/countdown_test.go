package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

	t.Run("sleep before every write", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assertOperationsEqual(t, spySleepPrinter.Calls, expected)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}

func assertStringEquals(t *testing.T, output, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("Expected: %q\nActual Output: %q", expected, output)
	}
}

func assertSleeperCalls(t *testing.T, actualCalls, expectedCalls int) {
	t.Helper()
	if actualCalls != expectedCalls {
		t.Errorf("Expected %d sleep calls\nGot %d sleep calls", expectedCalls, actualCalls)
	}
}

func assertOperationsEqual(t *testing.T, output, expected []string) {
	t.Helper()
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %q order of operations.\nGot: %q", expected, output)
	}
}
