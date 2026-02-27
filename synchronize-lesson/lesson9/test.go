package lesson9

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomWord(lenString int) string {
	const symbols = "abcdefghigkp"
	chars := make([]byte, lenString)

	for i := 0; i < lenString; i++ {
		char := symbols[rand.Intn(len(symbols))]
		chars[i] = char
	}

	return string(chars)
}

func writer(wg *sync.WaitGroup, lock *sync.RWMutex, counter map[string]int, nWrites int) {
	defer wg.Done()

	for ; nWrites > 0; nWrites-- {
		word := randomWord(3)
		lock.Lock()
		counter[word]++
		time.Sleep(time.Millisecond)
		lock.Unlock()
	}
}

func reader(wg *sync.WaitGroup, lock *sync.RWMutex, counter map[string]int, nReads int) {
	defer wg.Done()
	for ; nReads > 0; nReads-- {
		word := randomWord(3)
		lock.RLock()
		_ = counter[word]
		time.Sleep(time.Millisecond)
		lock.RUnlock()
	}
}

func Test() {
	counter := map[string]int{}
	var wg sync.WaitGroup
	wg.Add(5)
	var lock sync.RWMutex
	start := time.Now()

	go writer(&wg, &lock, counter, 100)
	go reader(&wg, &lock, counter, 100)
	go reader(&wg, &lock, counter, 100)
	go reader(&wg, &lock, counter, 100)
	go reader(&wg, &lock, counter, 100)

	wg.Wait()
	fmt.Println("Took", time.Since(start))
}
