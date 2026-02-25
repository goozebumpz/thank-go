package main

import "strings"

func countDigitsInWordsChan(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)
	stats := make(counter, 0)

	go func() {
		for _, word := range words {
			count := countDigits(word)
			counted <- count
		}
	}()

	for _, word := range words {
		stats[word] = <-counted
	}

	return stats
}
