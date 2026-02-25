package lesson6

import (
	"fmt"
	"time"
)

func TestTime() {
	work := func() {
		fmt.Println("work done")
	}

	var eventTime time.Time

	start := time.Now()
	timer := time.NewTimer(100 * time.Millisecond)

	go func() {
		eventTime = <-timer.C
		work()
	}()

	time.Sleep(200 * time.Millisecond)
	fmt.Printf("delayed function started %v\n", eventTime.Sub(start))
}

func TestTimeCancelled() {
	work := func() {
		fmt.Println("Work done")
	}
	start := time.Now()
	timer := time.NewTimer(100 * time.Millisecond)
	go func() {
		<-timer.C
		work()
	}()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("10ms has passed...")
	if timer.Stop() {
		fmt.Printf("delayed function canceled after %v \n", time.Since(start))
	}
}
