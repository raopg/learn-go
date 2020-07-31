package shapes

//Perimeter returns the perimeter of a rect given the length and breadth
func Perimeter(length, breadth float64) float64 {
	return 2 * (length + breadth)
}

//Area function returns the area of a given shape
func Area(length, breadth float64) float64 {
	return length * breadth
}
