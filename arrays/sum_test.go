package arrays

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {
	assertSumEquals := func(t *testing.T, output, expected int) {
		if output != expected {
			t.Errorf("Expected sum = %d\nActual sum = %d", expected, output)
		}
	}

	t.Run("Test sum function on slices - variable length arrays", func(t *testing.T) {

		testArray := []int{1, 2, 3}

		output := ArraySum(testArray)
		expected := 6

		assertSumEquals(t, output, expected)

	})
}

func TestSumAll(t *testing.T) {
	assertSumAllEquals := func(t *testing.T, output []int, expected []int) {
		if !reflect.DeepEqual(output, expected) { //Dangerous method to use because it is not typesafe
			t.Errorf("Expected sum = %d\nActual sum = %d", expected, output)
		}
	}

	output := SumAll([]int{1, 2, 3}, []int{1, 2, 3})
	expected := []int{6, 6}

	assertSumAllEquals(t, output, expected)
}
