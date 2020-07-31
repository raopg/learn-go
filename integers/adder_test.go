package integers

import "testing"

func TestAdder(t *testing.T) {
	assertAddition := func(t *testing.T, sum, expected int) {
		if sum != expected {
			t.Errorf("Expected sum = %d\nActual sum = %d", expected, sum)
		}
	}
	t.Run("Add two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertAddition(t, sum, expected)
	})
}
