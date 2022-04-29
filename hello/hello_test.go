package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Says hello to a person", func(t *testing.T) {
		got := Hello("Ben")
		want := "Hello, Ben"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Says hello to the world if a person is not provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}
