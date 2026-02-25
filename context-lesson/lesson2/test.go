package lesson2

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func execute(ctx context.Context, fn func() int) (int, error) {
	ch := make(chan int, 1)

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

func Test() {
	work := func() int {
		time.Sleep(200 * time.Millisecond)
		return 1
	}

	maybeCancel := func(cancel func()) {
		time.Sleep(100 * time.Millisecond)
		val := rand.Float32()
		fmt.Println(val)
		if val < 0.5 {
			cancel()
		}
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go maybeCancel(cancel)

	res, err := execute(ctx, work)
	fmt.Println(res, err)
}
