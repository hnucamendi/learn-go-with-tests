package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("in english", func(t *testing.T) {
		msg := Config(EnglishLang)
		got := msg.Hello("COATL")
		want := "Hello COATL!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		msg := Config(SpanishLang)
		got := msg.Hello("COATL")
		want := "Hola COATL!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		msg := Config(FrenchLang)
		got := msg.Hello("COATL")
		want := "Bonjour COATL!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in japanese", func(t *testing.T) {
		msg := Config(JapaneseLang)
		got := msg.Hello("COATL")
		want := "こんにちは COATL!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in other language", func(t *testing.T) {
		msg := Config(EnglishLang)
		got := msg.Hello("COATL")
		want := "Hello COATL!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("no parameter passed in", func(t *testing.T) {
		msg := Config(EnglishLang)
		got := msg.Hello("")
		want := "Hello World!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
