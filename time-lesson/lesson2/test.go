package lesson2

import (
	"errors"
	"fmt"
	"time"
)

func withWorkersError(n int, fn func()) (func() error, func()) {
	free := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		free <- struct{}{}
	}

	handle := func() error {
		select {
		case <-free:
			go func() {
				fn()
				free <- struct{}{}
			}()
			return nil
		default:
			return errors.New("Busy")
		}
	}

	wait := func() {
		for i := 0; i < n; i++ {
			<-free
		}
	}

	return handle, wait
}

func TestWithWorkersError() {
	fn := func() {
		time.Sleep(100 * time.Millisecond)
	}

	handle, wait := withWorkersError(2, fn)
	err := handle()
	fmt.Println("err1:", err)
	err = handle()
	fmt.Println("err2:", err)
	err = handle()
	fmt.Println("err3:", err)
	err = handle()
	fmt.Println("err4:", err)

	wait()
}
