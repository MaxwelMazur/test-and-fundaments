package hello

import "testing"

func TestHello(t *testing.T) {
	checkCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := hello("Max", "")
		want := "Hello, Max"
		checkCorrectMessage(t, got, want)
	})
	t.Run("says 'Hello world' when an empty string is passed", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, world"
		checkCorrectMessage(t, got, want)
	})
	t.Run("says 'Hallo' language german", func(t *testing.T) {
		got := hello("Max", "german")
		want := "Hallo, Max"
		checkCorrectMessage(t, got, want)
	})
}
