package math

import "testing"

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubtract(t *testing.T) {

	got := Subtract(40, 7)
	want := 33

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
