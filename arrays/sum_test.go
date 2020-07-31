package arrays

import "testing"

func TestArraySum(t *testing.T) {
	assertSumEquals := func(t *testing.T, output, expected int) {
		if output != expected {
			t.Errorf("Expected sum = %d\nActual sum = %d", expected, output)
		}
	}
	t.Run("Test sum function on arrays", func(t *testing.T) {

		testArray := []int{1, 2, 3, 4, 5}

		output := ArraySum(testArray)
		expected := 15

		assertSumEquals(t, output, expected)
	})

	t.Run("Test sum function on slices - variable length arrays", func(t *testing.T) {

		testArray := []int{1, 2, 3}

		output := ArraySum(testArray)
		expected := 6

		assertSumEquals(t, output, expected)

	})
}
