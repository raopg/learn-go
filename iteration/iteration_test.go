package iteration

import "testing"

func TestIteration(t *testing.T) {
	assertIterationSuccess := func(t *testing.T, output, expected string) {
		if output != expected {
			t.Errorf("Expected output: %q\nActual output: %q", output, expected)
		}
	}

	t.Run("Test repeated character string formation", func(t *testing.T) {
		output := Repeat("a", 5)
		expected := "aaaaa"

		assertIterationSuccess(t, output, expected)
	})

	t.Run("testing with multiple values", func(t *testing.T) {
		output := Repeat("abc", 2)
		expected := "abcabc"

		assertIterationSuccess(t, output, expected)
	})
}
