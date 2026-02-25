package lesson1

import (
	"fmt"
	"sync"
)

func TestWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("worker 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("worker 2")
	}()

	wg.Wait()
}

func TestWgGo() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println("Worker 1")
	})

	wg.Go(func() {
		fmt.Println("Worker 2")
	})

	wg.Wait()
}

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Work done")
}

func TestLinkWg() {
	var wg sync.WaitGroup
	wg.Add(2)
	work(&wg)
	work(&wg)
	wg.Wait()
	fmt.Println("All done")
}
