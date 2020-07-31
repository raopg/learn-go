package arrays

import "testing"

func TestArraySum(t *testing.T) {
	assertSumEquals := func(t *testing.T, output, expected int) {
		if output != expected {
			t.Errorf("Expected sum = %d\nActual sum = %d", expected, output)
		}
	}
	t.Run("Test sum function on arrays", func(t *testing.T) {

		testArray := [5]int{1, 2, 3, 4, 5}

		output := ArraySum(testArray)
		expected := 15

		assertSumEquals(t, output, expected)
	})
}
