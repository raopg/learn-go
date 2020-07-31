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
		{shape: Rectangle{breadth: 10, length: 5}, expected: 50.0},
		{shape: Circle{radius: 10}, expected: 314.1592653589793},
		{shape: Triangle{base: 10, height: 5}, expected: 25.0},
	}

	for _, test := range areaTest {
		output := test.shape.Area()
		if output != test.expected {
			t.Errorf("Expected = %g\nActual Output = %g", test.expected, output)
		}
	}
}
