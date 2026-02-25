package main

import (
	"fmt"
	"time"
)

func generateTestLockGor(cancel <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func modify(cancel <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	multiply := func(num int) int {
		time.Sleep(10 * time.Millisecond)
		return num * 2
	}

	go func() {
		defer fmt.Println("modify done")
		defer close(out)
		for num := range in {
			select {
			case out <- multiply(num):
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func print(in <-chan int) {
	for i := 0; i < 10; i++ {
		<-in
		fmt.Printf(".")
	}
	fmt.Println()
}

func testLockGor() {
	cancel := make(chan struct{})
	gen := generateTestLockGor(cancel)
	out := modify(cancel, gen)
	print(out)
	close(cancel)
	time.Sleep(50 * time.Millisecond)
}
