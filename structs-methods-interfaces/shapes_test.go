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

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		output := shape.Area()

		if output != expected {
			t.Errorf("Expected = %g\nActual Output = %g", expected, output)
		}
	}

	t.Run("test rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		checkArea(t, rectangle, 50.0)
	})

	t.Run("test circle area", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
