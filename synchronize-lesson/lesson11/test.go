package lesson11

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Test() {
	var wg sync.WaitGroup
	total := 0

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				total++
			}
		}()
	}

	wg.Wait()
	fmt.Println("total", total)
}

func Test2() {
	var n atomic.Int32
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				n.Add(1)

			}
		}()
	}

	wg.Wait()
	fmt.Println("total", n.Load())
}
