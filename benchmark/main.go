package main

import (
	"regexp"
	"strings"
)

type Counter map[string]int

var splitter *regexp.Regexp = regexp.MustCompile("n")

func WordCountRegexp(s string) Counter {
	counter := make(Counter)

	for _, word := range splitter.Split(s, -1) {
		word = strings.ToLower(word)
		counter[word]++
	}

	return counter
}
