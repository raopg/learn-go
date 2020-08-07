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
		{
			"Passing slices",
			[]Profile{
				{33, "Adam"},
				{31, "Eve"},
			},
			[]string{"Adam", "Eve"},
		},
		{
			"Passing arrays",
			[2]Profile{
				{33, "Adam"},
				{31, "Eve"},
			},
			[]string{"Adam", "Eve"},
		},
		{
			"Passing maps",
			map[string]string{
				"a": "b",
				"c": "d",
			},
			[]string{"b", "d"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"a": "b",
			"c": "d",
		}

		var output []string
		walk(aMap, func(input string) {
			output = append(output, input)
		})

		assertContains(t, output, "d")
		assertContains(t, output, "b")

	})
}

func assertContains(t *testing.T, output []string, expected string) {
	t.Helper()
	contains := false

	for _, x := range output {
		if x == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("Expected %+v to contain %q but did not find it", output, expected)
	}
}
