package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const src = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
const pattern = "commodo"

func BenchmarkMatchContains(b *testing.B) {
	for b.Loop() {
		MatchContains(src, pattern)
	}
}

func BenchmarkMatchRegexp(b *testing.B) {
	for b.Loop() {
		MatchRegexp(src, pattern)
	}
}

type IntSet struct {
	elems *[]int
}

func (is *IntSet) Contains(elem int) bool {
	for _, val := range *is.elems {
		if elem == val {
			return true
		}
	}

	return false
}

func (is *IntSet) Add(elem int) bool {
	if is.Contains(elem) {
		return false
	}

	*is.elems = append(*is.elems, elem)
	return true
}

func MakeIntSet() IntSet {
	return IntSet{elems: &[]int{}}
}

func randomSet(size int) IntSet {
	set := MakeIntSet()

	for i := 0; i < size; i++ {
		n := rand.Intn(1000000)
		set.Add(n)
	}

	return set
}

func BenchmarkIntSet(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000} {
		set := randomSet(size)
		name := fmt.Sprintf("Contains-%d", size)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				elem := rand.Intn(100000)
				set.Contains(elem)
			}
		})

	}
}
