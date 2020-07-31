package shapes

import "math"

//Shape interface contains common functions for all shapes
type Shape interface {
	Area() float64
}

//Rectangle struct defines the data-type for a rectangle's measurements
type Rectangle struct {
	length  float64
	breadth float64
}

//Circle struct defines the data-type for a circle's measurements
type Circle struct {
	radius float64
}

//Triangle struct defined the data-type for a triangle's measurements
type Triangle struct {
	base   float64
	height float64
}

//Perimeter returns the perimeter of a rect given the length and breadth
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.length + rect.breadth)
}

// This is the syntax for methods.

//Area provides the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.breadth * r.length
}

//Area provides the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//Area (Triangle) provides the area of a triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
