package lesson5

import (
	"context"
	"fmt"
	"time"
)

func execute(ctx context.Context, fn func() int) (int, error) {
	ch := make(chan int)

	go func() {
		ch <- fn()
	}()

	select {
	case val := <-ch:
		return val, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func TestSlowFast() {
	fast := func() int {
		time.Sleep(100 * time.Millisecond)
		return 42
	}

	slow := func() int {
		time.Sleep(300 * time.Millisecond)
		return 13
	}

	getDefaultCtx := func() (context.Context, context.CancelFunc) {
		const timeout = 200 * time.Millisecond
		return context.WithTimeout(context.Background(), timeout)
	}

	parentCtx, cancel := getDefaultCtx()
	defer cancel()
	childCtx, cancel := context.WithTimeout(parentCtx, 50*time.Millisecond)
	res, err := execute(childCtx, fast)
	defer cancel()
	fmt.Println(res, err)
	res, err = execute(parentCtx, slow)
	fmt.Println(res, err)
}
