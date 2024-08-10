package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("COATL")
	want := "Hello COATL!"

	if got != want {
		t.Errorf("got: %q, exppected: %q", got, want)
	}
}
