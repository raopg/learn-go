package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}
	t.Run("search a known word", func(t *testing.T) {

		output, _ := dictionary.Search("test")
		expected := "This is a test"

		assertStringEquals(t, output, expected)

	})

	t.Run("search an unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")
		expected := "Could not locate the word in dictionary"

		if err == nil {
			t.Fatal("Expected error, did not get one")
		}

		assertStringEquals(t, err.Error(), expected)
	})
}

func assertStringEquals(t *testing.T, output string, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("Expected: %s\nActual Output: %s", expected, output)
	}
}
