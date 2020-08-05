package reflection

import "testing"

func TestWalk(t *testing.T) {
	t.Run("test walk function", func(t *testing.T) {

		expected := "Prateek"

		var output []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			output = append(output, input)
		})

		if len(output) != 1 {
			t.Errorf("Expected 1 function call\nGot %d function calls", len(output))
		} else if output[0] != expected {
			t.Errorf("Expected %s\nGot %s", expected, output[0])
		}

	})
}
