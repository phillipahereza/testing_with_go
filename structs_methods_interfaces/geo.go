package geometry

import (
	"math"
)

// Rectangle 4 sided geometry
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle does
type Circle struct {
	Radius float64
}

// Triangle does
type Triangle struct {
	Height float64
	Base   float64
}

// Perimeter returns the perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns area
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

// Area returns area
func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Base * triangle.Height
}

// Area returns area
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

// Shape does shape things
type Shape interface {
	Area() float64
}
