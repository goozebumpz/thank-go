package lesson5

import (
	"fmt"
	"strconv"
	"strings"
)

func calcDistance(directions []string) int {
	var result float64 = 0

	for _, field := range directions {
		for _, word := range strings.Fields(field) {
			if strings.HasSuffix(word, "km") {
				value, err := strconv.ParseFloat(strings.Trim(word, "km"), 64)

				if err != nil {
					continue
				}
				result += value * 1000
				continue
			}

			if strings.HasSuffix(word, "m") {
				value, err := strconv.ParseFloat(strings.Trim(word, "m"), 64)
				if err != nil {
					continue
				}
				result += value
				continue
			}
		}
	}

	return int(result)
}

func Test() {
	text := "100m straight"
	res := calcDistance(strings.Fields(text))
	fmt.Println(res)
}
