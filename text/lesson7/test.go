package lesson7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func prettify(m map[string]int) string {
	if len(m) == 0 {
		return "{}"
	}

	if len(m) == 1 {
		var result string

		for key, value := range m {
			result = fmt.Sprintf("{ %s: %d }", key, value)
		}

		return result
	}

	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	slices.Sort(keys)
	b := strings.Builder{}
	b.WriteString("{\n")

	for _, word := range keys {
		b.WriteString("    ")
		b.WriteString(word)
		b.WriteString(": ")
		b.WriteString(strconv.Itoa(m[word]))
		b.WriteString(",")
		b.WriteString("\n")
	}
	b.WriteString("}")

	return b.String()
}

func Test() {
	//m := map[string]int{"one": 1, "two": 2, "three": 3}
	m := map[string]int{"one": 1}
	fmt.Println(prettify(m))
}
