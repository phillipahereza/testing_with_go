package geometry

// Rectangle 4 sided geometry
type Rectangle struct {
	Width float64
	Height float64
}


// Perimeter returns the perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns area
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}