package lesson2

import (
	"fmt"
	"math/rand"
	"sync"
)

func generate(nWords, wordLen int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for ; nWords > 0; nWords-- {
			out <- randomWord(wordLen)
		}
	}()

	return out
}

func randomWord(n int) string {
	const vowels = "eaiou"
	const consonants = "rtnslcdpm"
	chars := make([]byte, n)
	for i := 0; i < n; i += 2 {
		chars[i] = consonants[rand.Intn(len(consonants))]
	}
	for i := 1; i < n; i += 2 {
		chars[i] = vowels[rand.Intn(len(vowels))]
	}
	return string(chars)
}

func count(wg *sync.WaitGroup, in <-chan string, counter map[string]int) {
	defer wg.Done()

	for word := range in {
		counter[word]++
	}
}

func Test() {
	rand.Seed(0)

	in := generate(100, 3)
	counter := map[string]int{}

	var wg sync.WaitGroup
	wg.Add(2)

	go count(&wg, in, counter)
	go count(&wg, in, counter)

	wg.Wait()

	fmt.Println(counter)
}
