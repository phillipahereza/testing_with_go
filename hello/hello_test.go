package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be
		// in our function call rather than inside our test helper
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s' ", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := hello("Phillip", "")
		want := "Hello, Phillip"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := hello("Payet", "French")
		want := "Bonjour, Payet"

		assertCorrectMessage(t, got, want)
	})

}
