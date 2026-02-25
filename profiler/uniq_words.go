package main

import (
	"cmp"
)

func UniqWords(str string) []string {
	words := splitString(str)
	words = sortWords(words)
	words = uniqWords(words)

	return words
}

func splitString(s string) []string {
	words := []string{}
	runes := []rune{}

	for _, r := range s {
		if r != ' ' {
			runes = append(runes, r)
			continue
		}

		if r == ' ' && len(runes) > 0 {
			words = append(words, string(runes))
			runes = []rune{}
		}
	}

	if len(runes) > 0 {
		words = append(words, string(runes))
	}

	return words
}

func sortWords(l []string) []string {
	return fastSorting(l)
}

func uniqWords(l []string) []string {
	cash := map[string]bool{}
	result := []string{}

	for _, word := range l {
		if _, has := cash[word]; has {
			continue
		}
		cash[word] = true
		result = append(result, word)
	}

	return result
}

func fastSorting[T cmp.Ordered](list []T) []T {
	if len(list) <= 1 {
		return list
	}

	greater := []T{}
	less := []T{}
	equal := []T{}
	targetIndex := len(list) / 2
	targetValue := list[targetIndex]

	for _, item := range list {
		if item < targetValue {
			less = append(less, item)
		} else if item > targetValue {
			greater = append(greater, item)
		} else {
			equal = append(equal, item)
		}
	}

	var result = append(fastSorting(less), equal...)
	result = append(result, fastSorting(greater)...)

	return result
}
