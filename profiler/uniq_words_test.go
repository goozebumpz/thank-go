package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func TestUniqWords(t *testing.T) {
	tests := map[string][]string{
		"in a coat of gold or  a coat of red a lion still has claws": {"a", "claws", "coat", "gold", "has", "in", "lion", "of", "or", "red", "still"},
		"hello world":   {"hello", "world"},
		" hello world ": {"hello", "world"},
		"world hello":   {"hello", "world"},
		"hello":         {"hello"},
		" ":             {},
		"":              {},
	}

	for phrase, want := range tests {
		got := UniqWords(phrase)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("'%s': got %v, want %v", phrase, got, want)
		}
	}
}

func BenchmarkUniqWords(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000, 50000} {
		name := fmt.Sprintf("UniqWords-%d", size)
		phrase := randomPhrase(size)

		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				UniqWords(phrase)
			}
		})
	}
}

func randomWord(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func randomPhrase(n int) string {
	words := make([]string, n)

	for i := range words {
		words[i] = randomWord(n)
	}

	return strings.Join(words, " ")
}
