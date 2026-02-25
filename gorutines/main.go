package main

import (
	"fmt"
	"strings"
)

type nextFunc func() string

func random(phrase string) nextFunc {
	words := strings.Split(phrase, " ")
	idx := 0

	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]

		idx++
		return word
	}
}

func main() {
	next := random("welcome to the secret shop")

	for word := next(); word != ""; word = next() {
		fmt.Println(word)
	}
}
