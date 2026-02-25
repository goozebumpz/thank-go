package main

import (
	"fmt"
	"testing"
)

func TestIntMin(t *testing.T) {
	got := IntMin(2, -2)
	want := -2

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestIntMinDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, test := range tests {
		name := fmt.Sprintf("case(%d,%d)", test.a, test.b)

		t.Run(name, func(t *testing.T) {
			got := IntMin(test.a, test.b)

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
