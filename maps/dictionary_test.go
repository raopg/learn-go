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
		expected := ErrNotFound

		if err == nil {
			t.Fatal("Expected error, did not get one")
		}

		assertErrorEquals(t, err, expected)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add a new word to dictionary", func(t *testing.T) {
		d := Dictionary{}
		d.Add("test", "this is just a test")

		expected := "this is just a test"

		output, err := d.Search("test")

		if err != nil {
			t.Fatal("Found an error when expected none while adding a valid entry")
		}

		assertStringEquals(t, output, expected)
	})
}

func assertStringEquals(t *testing.T, output string, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("Expected: %s\nActual Output: %s", expected, output)
	}
}

func assertErrorEquals(t *testing.T, err error, expected error) {
	t.Helper()

	if err != expected {
		t.Errorf("Expected = %q\n Actual error = %q\n", expected, err)
	}
}
