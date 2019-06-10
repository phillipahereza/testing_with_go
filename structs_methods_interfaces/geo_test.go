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

	t.Run("rectangles", func(t *testing.T){
		rectangle := Rectangle{20.0, 4.0}
		got := rectangle.Area()
		wanted := 80.0
	
		if got != wanted {
			t.Fatalf("got %.2f but wanted %.2f", got, wanted)
		}
	})

	t.Run("circles", func(t *testing.T){
		circle := Circle{10.0}
		got := circle.Area()
		wanted := 314.1592653589793

		if got != wanted {
			t.Fatalf("got %.2f but wanted %.2f", got, wanted)
		}
	})
	
}