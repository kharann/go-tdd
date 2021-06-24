package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kharann")
	want := "Hello, Kharann"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
