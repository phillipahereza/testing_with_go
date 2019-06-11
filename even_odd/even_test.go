package main

import (
	"testing"
)

func TestEven(t *testing.T) {

	t.Run("test 2 is even", func(t *testing.T) {
		got := isEvenOdd(2)
		wanted := "even"

		if got != wanted {
			t.Fatalf("Got %s but expected %s", got, wanted)
		}
	})

	t.Run("test 3 is odd", func(t *testing.T) {
		got := isEvenOdd(3)
		wanted := "odd"

		if got != wanted {
			t.Fatalf("Got %s but expected %s", got, wanted)
		}
	})
	
}