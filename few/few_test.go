package main

import (
	"fmt"
	"testing"
)

func short(t *testing.T, message string) {
	if testing.Short() {
		t.Skip(message)
	}
}

func TestSumFew(t *testing.T) {
	if Sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}

func TestSumN(t *testing.T) {
	short(t, "skipping test in short mode.")
	n := 1_000_000_000
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	got := Sum(nums...)
	want := n * (n + 1) / 2

	if got != want {
		t.Errorf("Expected sum[i=1..n](i) == n*(n+1)/2")
	}
}
func TestMain(m *testing.M) {
	fmt.Println("Testing Sum(n1, n2, ...,nk)...")
	fmt.Println("Finished testing")
}

func TestSumZero(t *testing.T) {
	if Sum() != 0 {
		t.Errorf("Expected Sum() == 0")
	}
}

func TestSumOne(t *testing.T) {
	if Sum(1) != 1 {
		t.Errorf("Expected Sum(1) == 1")
	}
}

func TestSumPair(t *testing.T) {
	t.Skip()
	if Sum(1, 2) != 3 {
		t.Errorf("Expected Sum(1, 2) == 3")
	}
}

func TestSumMany(t *testing.T) {
	if Sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}
