package main

import (
	"fmt"
	"time"
	"time-lesson/lesson11"
)

func main() {
	work := func() {
		fmt.Println("work")
	}

	start := time.Now()
	handle, cancel := lesson11.Throttle2(5, work)
	defer cancel()

	handle()
	handle()
	handle()
	handle()
	handle()
	handle()
	handle()
	handle()
	handle()
	handle()

	fmt.Printf("%dms \n", time.Since(start).Milliseconds())
}
