package lesson4

import (
	"fmt"
	"time"
)

func delay(duration time.Duration, fn func()) func() {
	canceled := false

	go func() {
		time.Sleep(duration)
		if !canceled {
			fn()
		}
	}()

	cancel := func() {
		canceled = true
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
