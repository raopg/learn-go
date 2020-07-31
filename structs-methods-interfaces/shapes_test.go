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
	t.Run("test rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		output := rectangle.Area()
		expected := 50.0

		assertShapeDataEquals(t, output, expected)
	})

	t.Run("test circle area", func(t *testing.T) {
		circle := Circle{10.0}
		output := circle.Area()
		expected := 314.1592

		assertShapeDataEquals(t, output, expected)
	})
}
