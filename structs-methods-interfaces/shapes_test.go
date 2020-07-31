package shapes

import "testing"

func assertShapeDataEquals(t *testing.T, output, expected float64) {
	t.Helper()

	if output != expected {
		t.Errorf("Expected = %.2f\nActual = %.2f", expected, output)
	}
}

func TestPerimeter(t *testing.T) {
	output := Perimeter(10.0, 10.0)
	expected := 40.0

	assertShapeDataEquals(t, output, expected)
}

func TestArea(t *testing.T) {
	output := Area(10.0, 5.0)
	expected := 50.0

	assertShapeDataEquals(t, output, expected)
}
