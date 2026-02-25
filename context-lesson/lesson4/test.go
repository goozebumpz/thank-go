package lesson4

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

func TestTimeoutContext() {
	work := func() int {
		time.Sleep(100 * time.Millisecond)
		return 1
	}

	randomChoice := func(arg ...int) int {
		i := rand.Intn(len(arg))
		return arg[i]
	}

	timeout := time.Duration(randomChoice(50, 150)) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := execute(ctx, work)
	fmt.Println(res, err)
}
