package lesson6

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

func TestDeadline() {
	work := func() int {
		time.Sleep(100 * time.Millisecond)
		return 42
	}

	randomChoice := func(args ...int) int {
		i := rand.Intn(len(args))
		return args[i]
	}

	timeout := time.Duration(randomChoice(50, 150)) * time.Millisecond
	deadline := time.Now().Add(timeout)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	deadline, ok := ctx.Deadline()
	fmt.Println(deadline, ok)
	defer cancel()

	res, err := execute(ctx, work)
	fmt.Println(res, err)
}
