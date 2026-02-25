package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func randomPhrase(n int) string {
	words := make([]string, n)

	for i := range words {
		words[i] = randomWord(3)
	}

	return strings.Join(words, " ")
}

func randomWord(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func BenchmarkRegexp(b *testing.B) {
	for _, lenght := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(lenght)
		name := fmt.Sprintf("Regexp-%d", lenght)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountRegexp(phrase)
			}
		})
	}

}

func BenchmarkFields(b *testing.B) {
	for _, lenght := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(lenght)
		name := fmt.Sprintf("Regexp-%d", lenght)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountFields(phrase)
			}
		})
	}

}

func BenchmarkSplit(b *testing.B) {
	for _, lenght := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(lenght)
		name := fmt.Sprintf("Regexp-%d", lenght)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountSplit(phrase)
			}
		})
	}
}

func BenchmarkLowerCase(b *testing.B) {
	for _, lenght := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(lenght)
		name := fmt.Sprintf("Regexp-%d", lenght)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountLowerPhrase(phrase)
			}
		})
	}
}

func BenchmarkCountAllocate(b *testing.B) {
	for _, lenght := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(lenght)
		name := fmt.Sprintf("Regexp-%d", lenght)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountAllocate(phrase)
			}
		})
	}
}
