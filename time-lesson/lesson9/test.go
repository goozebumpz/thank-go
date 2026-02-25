package lesson9

import (
	"fmt"
	"time"
)

func TestTicker() {
	fn := func() {
		fmt.Println("What's up bitch")
	}
	done := make(chan struct{})

	ticker := time.NewTicker(time.Second)

	go func() {
		counter := 0
		for counter < 10 {
			tick := <-ticker.C
			fmt.Println("Tick", tick)
			fn()
			counter++
		}
		ticker.Stop()
		done <- struct{}{}
	}()

	<-done
}

func TestTickerStop() {
	work := func(at time.Time) {
		fmt.Println(at)
		time.Sleep(100 * time.Millisecond)
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for {
			at := <-ticker.C
			work(at)
		}
	}()

	time.Sleep(260 * time.Millisecond)
}
