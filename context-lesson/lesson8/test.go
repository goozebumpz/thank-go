package lesson8

import (
	"context"
	"errors"
	"fmt"
)

func Test() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	fmt.Println(errors.Is(ctx.Err(), context.Canceled))
	fmt.Println(ctx.Err())
}

func TestCancelCause() {
	ctx, cancel := context.WithCancelCause(context.Background())
	err := errors.New("because you suck")

	cancel(err)
	fmt.Println(ctx.Err())
	fmt.Println(context.Cause(ctx))
}
