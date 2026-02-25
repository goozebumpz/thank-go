// Сколько цифр в каждом слове?
package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type counter map[string]int

// начало решения

// countDigitsInWords считает количество цифр в словах фразы.
func countDigitsInWords(phrase string) counter {
	var wg sync.WaitGroup
	syncStats := new(sync.Map)
	words := strings.Fields(phrase)

	for _, word := range words {
		wg.Add(1)

		go func() {
			defer wg.Done()
			count := countDigits(word)
			syncStats.Store(word, count)
		}()
	}

	wg.Wait()

	return asStats(syncStats)
}

// конец решения

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func testCountDigits() {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	printStats(counts)
}
