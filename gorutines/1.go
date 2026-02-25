package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func say(worker int, str string) {
	for _, s := range strings.Fields(str) {
		fmt.Printf("Worker %d say some: %s... \n", worker, s)
		dur := time.Duration(rand.Intn(200)) * time.Millisecond
		time.Sleep(dur)
	}
}

func test1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		say(1, "Bitch stupid suck my dick")
	}()

	go func() {
		defer wg.Done()
		say(2, "Bitch stupid suck my dick")
	}()

	wg.Wait()
}
