package shapes

import "testing"

func assertShapeDataEquals(t *testing.T, output, expected float64) {
	t.Helper()

	if output != expected {
		t.Errorf("Expected = %g\nActual = %g", expected, output)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	output := Perimeter(rectangle)
	expected := 40.0

	assertShapeDataEquals(t, output, expected)
}

func TestArea(t *testing.T) {
	// table-driven tests => where tables of data are used to test a function
	areaTest := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{5, 10}, 25.0},
	}

	for _, test := range areaTest {
		output := test.shape.Area()
		if output != test.expected {
			t.Errorf("Expected = %g\nActual Output = %g", test.expected, output)
		}
	}
}
