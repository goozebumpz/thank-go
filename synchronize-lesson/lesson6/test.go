package lesson6

import (
	"fmt"
	"math/rand"
	"sync"
)

func generateWord(n int) string {
	const symbolsA = "asdfghjkl"
	chars := make([]byte, n)

	for i := 0; i < n; i++ {
		chars[i] = symbolsA[rand.Intn(len(symbolsA))]
	}

	return string(chars)
}

func generate(count, lenWord int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- generateWord(lenWord)
		}
	}()

	return out
}

func count(wg *sync.WaitGroup, in <-chan string, counters []map[string]int, idx int) {
	defer wg.Done()
	counter := map[string]int{}
	for word := range in {
		counter[word]++
	}
	counters[idx] = counter
}

func merge(counters ...map[string]int) map[string]int {
	merged := map[string]int{}

	for _, counter := range counters {
		for word, freq := range counter {
			merged[word] += freq
		}
	}

	return merged
}

func Test() {
	in := generate(100, 3)
	counters := make([]map[string]int, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go count(&wg, in, counters, 0)
	go count(&wg, in, counters, 1)

	wg.Wait()
	result := merge(counters...)
	fmt.Println(result)
}
