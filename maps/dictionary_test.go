package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search a word", func(t *testing.T) {

		dictionary := Dictionary{"test": "This is a test"}

		output := Search(dictionary, "test")
		expected := "This is a test"

		assertSearchEquals(t, output, expected)

	})
}

func assertSearchEquals(t *testing.T, output string, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("Expected: %s\nActual Output: %s", expected, output)
	}
}
