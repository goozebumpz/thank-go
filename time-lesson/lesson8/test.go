package lesson8

import (
	"fmt"
	"math/rand"
	"time"
)

func ResetTimer() {
	nums := make(chan int)
	done := make(chan struct{})

	go func() {
		for n := range nums {
			fmt.Print(n)
		}
		close(done)
	}()

	for i := range 10 {
		delay := time.Duration(5+rand.Intn(11)) * time.Millisecond
		time.Sleep(delay)
		nums <- i
	}

	close(nums)
	<-done
}

func ResetTimer10Seconds() {
	nums := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(done)
		timeout := 10 * time.Millisecond
		timer := time.NewTimer(timeout)

		for {
			select {
			case <-nums:
			case <-timer.C:
				fmt.Print("!")
			}
			timer.Reset(timeout)
		}
	}()

	for i := range 10 {
		delay := time.Duration(5+rand.Intn(11)) * time.Millisecond
		time.Sleep(delay)
		nums <- i
	}

	close(nums)
	<-done
}

func TimerResetAfterFunc() {
	var start time.Time

	work := func() {
		fmt.Printf("work done after %dms\n", time.Since(start).Milliseconds())
	}

	timeout := time.Millisecond / 12
	start = time.Now()
	timer := time.AfterFunc(timeout, work)

	delay := time.Duration(rand.Intn(11)+5) * time.Millisecond
	time.Sleep(delay)
	fmt.Printf("%dms has passed... \n", delay.Milliseconds())

	timer.Reset(timeout)
	start = time.Now()
	time.Sleep(50 * time.Millisecond)
}

func TestAfterFunc() {
	fn := func() {
		fmt.Println("Work hard")
	}
	start := time.Now()
	fmt.Printf("Start 1: %dms", time.Since(start).Milliseconds())
	timeout := 100 * time.Millisecond

	timer := time.AfterFunc(timeout, fn)

	timer.Reset(timeout)

	time.Sleep(1000 * time.Millisecond)
}
