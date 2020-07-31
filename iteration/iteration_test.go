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

//BenchmarkRepeat is a benchmark test for the repeat func
func BenchmarkRepeat(b *testing.B) {
	// Since our function scales on the times variable, we will limit that
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
