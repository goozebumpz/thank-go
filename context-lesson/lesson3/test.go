package lesson3

import (
	"context"
	"fmt"
)

func generate(ctx context.Context, start int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := start; ; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func TestGenerate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generated := generate(ctx, 11)

	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			break
		}
	}
	fmt.Println()
}
