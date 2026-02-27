package lesson3

import (
	"fmt"
	"sync"
)

func Test() {
	var total int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		total++
	}()

	go func() {
		defer wg.Done()
		total++
	}()

	wg.Wait()
	fmt.Println(total)
}
