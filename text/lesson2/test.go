package lesson2

import (
	"fmt"
	"strings"
)

func slugify(src string) string {
	const symbols = "qwertyuiopasdfghjklzxcvbnm1234567890"
	runes := []rune(strings.ToLower(src))
	result := []string{}
	candidate := []rune{}

	for _, r := range runes {
		isSafe := strings.Contains(symbols, string(r)) || r == '-'

		if isSafe {
			candidate = append(candidate, r)
		} else if !isSafe && len(candidate) > 0 {
			result = append(result, string(candidate))
			candidate = []rune{}
		} else {
			candidate = []rune{}
		}
	}

	if len(candidate) > 0 {
		result = append(result, string(candidate))
	}

	return strings.Join(result, "-")
}

func Test() {
	test := "!Attention, attention!"
	fmt.Println(slugify(test))
}
