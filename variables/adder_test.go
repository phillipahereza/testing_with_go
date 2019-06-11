package integers

import "testing"

func TestAdder(t *testing.T) {

	assertEquals := func(sum, expected int, t *testing.T) {
		t.Helper()
		if sum != expected {
			t.Errorf("Expected %d but got %d", expected, sum)
		}
	}

	t.Run("2+2=4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertEquals(sum, expected, t)

	})

	t.Run("1+6=7", func(t *testing.T) {
		sum := Add(1, 6)
		expected := 7

		assertEquals(sum, expected, t)
	})

}
