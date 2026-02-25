package main

import (
	"strings"
)

func countDigitsInWordGenerator(next nextFunc) counter {
	stats := counter{}

	for {
		word := next()
		if word == "" {
			break
		}
		count := countDigits(word)
		stats[word] = count
	}

	return counter{}
}

func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
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

func testGenerator() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsGen(next)
	printStats(stats)
}

type pair struct {
	word  string
	count int
}

func countDigitsInWordsGen1(next nextFunc) counter {
	stats := counter{}
	counted := make(chan pair)

	go func() {
		for {
			word := next()
			counted <- pair{word, countDigits(word)}

			if word == "" {
				break
			}

		}
	}()

	for {
		pair := <-counted

		if pair.word == "" {
			break
		}

		stats[pair.word] = pair.count
	}

	return stats
}

func countDigitsInWordsGen2(next nextFunc) counter {
	pending := make(chan string)
	counted := make(chan pair)
	stats := counter{}

	go func() {
		for {
			str := next()
			pending <- str
			if str == "" {
				break
			}
		}
	}()

	go func() {
		for {
			str := <-pending
			count := countDigits(str)
			counted <- pair{str, count}
			if str == "" {
				break
			}
		}
	}()

	for {
		pair := <-counted

		if pair.word == "" {
			break
		}

		stats[pair.word] = pair.count
	}

	return stats
}

func submitWords(ch chan string, next nextFunc) {
	for {
		str := next()
		ch <- str
		if str == "" {
			break
		}
	}
}

func countWords(charWord chan string, chanPair chan pair) {
	for {
		str := <-charWord
		count := countDigits(str)
		chanPair <- pair{str, count}
		if str == "" {
			break
		}
	}
}

func fillStats(chanPair chan pair) counter {
	stats := counter{}

	for {
		pair := <-chanPair

		if pair.word == "" {
			break
		}

		stats[pair.word] = pair.count
	}

	return stats
}

func countDigitsInWordsGen(next nextFunc) counter {
	chanStr := make(chan string)
	go submitWords(chanStr, next)
	chanPair := make(chan pair)
	go countWords(chanStr, chanPair)

	return fillStats(chanPair)
}

func submitWords1(next nextFunc, out chan string) {
	for word := next(); word != ""; word = next() {
		out <- word
	}
	close(out)
}

func countWords2(in chan string, out chan pair) {
	for word := range in {
		count := countDigits(word)
		pair := pair{word, count}
		out <- pair
	}
	close(out)
}

func fillStats2(in chan pair) counter {
	stats := counter{}

	for pair := range in {
		stats[pair.word] = pair.count
	}

	return stats
}
