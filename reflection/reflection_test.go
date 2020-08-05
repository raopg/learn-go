package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var output []string
			walk(test.Input, func(input string) {
				output = append(output, input)
			})

			if !reflect.DeepEqual(output, test.ExpectedCalls) {
				t.Errorf("Expected: %v, Got: %v", test.ExpectedCalls, output)
			}
		})
	}
}
