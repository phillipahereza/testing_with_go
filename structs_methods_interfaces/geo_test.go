package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 4.2}
	got := Perimeter(rectangle)
	wanted := 28.4

	if got != wanted {
		t.Fatalf("got %.2f but wanted %.2f", got, wanted)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 20.0, Height: 4.0}, hasArea: 80.00},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Height: 12.0, Base: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f but wanted %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}
}
