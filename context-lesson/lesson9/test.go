package lesson9

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	select {
	case <-time.After(100 * time.Millisecond):
	case <-ctx.Done():
	}
}

func cleanup() {
	fmt.Println("cleanup")
}

func TestWork() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)

	start := time.Now()
	work(ctx)

	if ctx.Err() != nil {
		fmt.Println("canceled after", time.Since(start).Milliseconds(), "ms")
	}
	context.AfterFunc(ctx, cleanup)
}

func RegisterAfterCancelContext() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	start := time.Now()
	work(ctx)

	if ctx.Err() != nil {
		fmt.Println("canceled after", time.Since(start))
	}
	context.AfterFunc(ctx, cleanup)
}

func CancelRegistration() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	stopCleanup := context.AfterFunc(ctx, cleanup)

	stopped := stopCleanup()
	work(ctx)

	fmt.Println("stopped cleanup:", stopped)
}
