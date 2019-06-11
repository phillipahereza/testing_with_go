package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected '%s' but got '%s'", expected, got)
	}
}

func ExampleRepeat() {
	repeated := Repeat("A", 2)
	fmt.Println(repeated)
	// Output: AA
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A", 23)
	}
}
