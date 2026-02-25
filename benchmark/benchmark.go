// Package benchmark
package main

import (
	"regexp"
	"strings"
)

func MatchContains(src, pattern string) bool {
	return strings.Contains(src, pattern)
}

func MatchRegexp(src, pattern string) bool {
	re, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return re.MatchString(src)
}
