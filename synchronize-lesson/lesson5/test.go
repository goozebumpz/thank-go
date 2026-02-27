package lesson5

import (
	"fmt"
	"time"
)

func delay(duration time.Duration, fn func()) func() {
	alive := make(chan struct{})
	close(alive)

	go func() {
		time.Sleep(duration)
		select {
		case <-alive:
			fn()
		default:
		}
	}()

	cancel := func() {
		alive = nil
	}

	return cancel
}

func Test() {
	work := func() {
		fmt.Println("Work done")
	}
	cancel := delay(50*time.Millisecond, work)
	time.Sleep(50 * time.Millisecond)
	go cancel()
}
