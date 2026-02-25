package main

import (
	"sort"
	"strings"
)

func main() {

}

func split(str string) []string {
	return strings.Fields(str)
}

func join(l1, l2 []string) []string {
	list := make([]string, len(l1)+len(l2))
	idx := 0

	for _, word := range l1 {
		list[idx] = word
		idx++
	}

	for _, word := range l2 {
		list[idx] = word
		idx++
	}

	return list
}

func lower(words []string) []string {
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

func sorted(words []string) []string {
	sort.Strings(words)
	return words
}

func uniq(words []string) []string {
	uniq := []string{}
	current := ""

	for _, word := range words {
		if word == current {
			continue
		}
		if current != "" {
			uniq = append(uniq, current)
		}
		current = word
	}

	if current != "" {
		uniq = append(uniq, current)
	}

	return uniq
}

func JoinWords(first, second string) []string {
	words1 := split(first)
	words2 := split(second)
	words := join(words1, words2)
	words = lower(words)
	words = sorted(words)
	words = uniq(words)
	return words
}
