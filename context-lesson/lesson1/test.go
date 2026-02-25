package lesson1

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func execute(cancel <-chan struct{}, fn func() int) (int, error) {
	ch := make(chan int, 1)

	go func() {
		ch <- fn()
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-cancel:
		return 0, errors.New("canceled")
	}
}

func executeContext(ctx context.Context, fn func() int) (int, error) {
	ch := make(chan int, 1)

	go func() {
		ch <- fn()
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func Test1() {
	work := func() int {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("work done")
		return 42
	}

	maybeCancel := func(cancel chan struct{}) {
		time.Sleep(50 * time.Millisecond)
		if rand.Float32() < 0.5 {
			close(cancel)
		}
	}

	cancel := make(chan struct{})
	go maybeCancel(cancel)
	res, err := execute(cancel, work)
	fmt.Println(res, err)
}
