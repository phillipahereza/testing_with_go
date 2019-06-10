package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 4.2)
	wanted := 28.4

	if got != wanted {
		t.Fatalf("got %.2f but wanted %.2f", got, wanted)
	}
}