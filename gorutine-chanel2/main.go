package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// func main() {
// 	done := make(chan struct{})
// 	phrases := []string{
// 		"go is awesome",
// 		"cats are cute",
// 		"rain is wet",
// 		"channels are hard",
// 		"floor is lava",
// 	}

// 	for idx, str := range phrases {
// 		go say(done, idx, str)
// 	}

// 	for i := 0; i < len(phrases); i++ {
// 		<-done
// 	}
// }

func say(done chan<- struct{}, id int, pharse string) {
	for _, word := range strings.Fields(pharse) {
		fmt.Printf("Worker %d says: %s... \n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	done <- struct{}{}
}

// func main() {
// 	var wg sync.WaitGroup

// 	phrases := []string{
// 		"go is awesome",
// 		"cats are cute",
// 		"rain is wet",
// 		"channels are hard",
// 		"floor is lava",
// 	}
// 	wg.Add(5)

// 	for idx, phrase := range phrases {
// 		go say(&wg, idx, phrase)
// 	}

// 	wg.Wait()
// }

// func say(wg *sync.WaitGroup, idx int, str string) {
// 	for _, word := range strings.Fields(str) {
// 		fmt.Printf("Gorutine %d say: %s... \n", idx, word)
// 	}
// 	wg.Done()
// }
