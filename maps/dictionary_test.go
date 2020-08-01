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
		word := "test"
		addError := d.Add(word, "this is just a test")

		expected := "this is just a test"

		assertDefinition(t, d, word, expected)
		assertErrorEquals(t, addError, nil)

	})

	t.Run("Add an existing  to dictionary", func(t *testing.T) {
		word := "test"
		definition := "this is a test"

		d := Dictionary{word: definition}
		err := d.Add("test", "this is just a test")

		assertErrorEquals(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update an existing value", func(t *testing.T) {
		word := "test"
		oldVal := "This is outdated; update me!"
		newVal := "Updated!"
		d := Dictionary{word: oldVal}

		d.Update(word, newVal)

		assertDefinition(t, d, word, newVal)

	})
	t.Run("Update a new value", func(t *testing.T) {
		word := "test"
		val := "Updated!"
		d := Dictionary{}

		err := d.Update(word, val)

		if err == nil {
			t.Fatal("Expected error, got none")
		}

		assertErrorEquals(t, err, ErrWordDoesNotExist)

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

func assertDefinition(t *testing.T, d Dictionary, word, expectedDefinition string) {
	t.Helper()

	definition, err := d.Search(word)

	if err != nil {
		t.Fatal("Did not expected an error but received one")
	}

	if definition != expectedDefinition {
		t.Errorf("Expected = %s\nActual definition = %s", expectedDefinition, definition)
	}
}
