package ints

import "testing"

func TestIntMin(t *testing.T) {
	got := IntMin(1, 2)
	want := 1

	if want != got {
		t.Errorf("Want %d not equal got %d", want, got)
	}
}

func TestIntMinEquals(t *testing.T) {
	got := IntMin(2, 2)
	want := 2

	if want != got {
		t.Errorf("Want %d not equal got %d", want, got)
	}
}

func TestIntMinFirst(t *testing.T) {
	got := IntMin(4, 2)
	want := 2

	if want != got {
		t.Errorf("Want %d not equal got %d", want, got)
	}
}
