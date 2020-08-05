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
			}{"Prateek"},
			[]string{"Prateek"},
		},
		{
			"Struct with 2/multiple string fields",
			struct {
				Name string
				City string
			}{"Prateek", "Los Angeles"},
			[]string{"Prateek", "Los Angeles"},
		},
		{
			"Struct with a non-string field",
			struct {
				Name string
				City string
				Age  int
			}{"Prateek", "Los Angeles", 21},
			[]string{"Prateek", "Los Angeles"},
		},
		{
			"Nested fields",
			Person{
				"Prateek",
				Profile{21, "London"},
			},
			[]string{"Prateek", "London"},
		},
		{
			"Passing a pointer to a struct",
			&Person{
				"Prateek",
				Profile{21, "London"},
			},
			[]string{"Prateek", "London"},
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
