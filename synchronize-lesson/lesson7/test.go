package lesson7

import (
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

func count(wg *sync.WaitGroup, lock *sync.Mutex, in <-chan string, counter map[string]int) {
	defer wg.Done()

	for word := range in {
		lock.Lock()
		counter[word]++
		lock.Unlock()
	}
}

func Test() {
	in := generate(100, 3)
	counter := map[string]int{}

	var wg sync.WaitGroup
	wg.Add(2)

	var lock sync.Mutex
	
	go count(&wg, &lock, in, counter)
	go count(&wg, &lock, in, counter)

	wg.Wait()
}
