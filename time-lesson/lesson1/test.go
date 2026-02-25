package lesson1

import (
	"fmt"
	"time"
)

func withWorkers(n int, fn func()) (handle func(), wait func()) {
	free := make(chan int, n)

	for i := 0; i < n; i++ {
		free <- n
	}

	handle = func() {
		<-free

		go func() {
			fn()
			free <- 0
		}()
	}

	wait = func() {
		for i := 0; i < n; i++ {
			<-free
			fmt.Println(i)
		}
	}

	return handle, wait
}

func Test() {
	timeNow := time.Now()
	work := func() {
		time.Sleep(50 * time.Millisecond)
	}

	work()
	work()
	work()
	work()
	fmt.Printf("Time is %d", time.Since(timeNow).Milliseconds())
}

func TestWithWorkers() {
	fn := func() {
		time.Sleep(100 * time.Millisecond)
	}

	start := time.Now()

	handle, wait := withWorkers(2, fn)

	handle()
	handle()
	handle()
	handle()
	wait()

	fmt.Printf("Time is %d ms", time.Since(start).Milliseconds())
}
