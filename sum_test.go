package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Zero", []int{1, 0, 1, -2}, 0},
		{"One", []int{0, 1, 1, -1}, 1},
		{"Two", []int{2, 0, 1, -1}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := Sum(test.nums...)

			if sum != test.want {
				t.Errorf("%s, want %d, sum %d", test.name, test.want, sum)
			}
		})
	}

}
