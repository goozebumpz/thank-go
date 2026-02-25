package lesson10

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var ErrFailed = errors.New("failed")
var ErrManual = errors.New("manual")

type Worker struct {
	fn        func() error
	ctx       context.Context
	ctxCancel context.CancelCauseFunc
	isStarted bool
}

func NewWorker(fn func() error) *Worker {
	ctx, cancel := context.WithCancelCause(context.Background())

	return &Worker{
		fn:        fn,
		ctx:       ctx,
		ctxCancel: cancel,
		isStarted: false,
	}
}

func (w *Worker) Start() {
	if w.isStarted {
		fmt.Println("Work")
		return
	}

	w.isStarted = true

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			default:
				err := w.fn()
				if err != nil {
					w.ctxCancel(ErrFailed)
					return
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	select {
	case <-w.ctx.Done():
	default:
		w.isStarted = false
		w.ctxCancel(ErrManual)
	}
}

func (w *Worker) AfterStop(fn func()) {
	if w.isStarted {
		return
	}

	select {
	case <-w.ctx.Done():
	default:
		context.AfterFunc(w.ctx, fn)
	}
}

func (w *Worker) Err() error {
	return context.Cause(w.ctx)
}

func Test() {
	count := 3
	fn := func() error {
		fmt.Print(count, " ")
		count--
		if count == 0 {
			return errors.New("count is zero")
		}
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	worker := NewWorker(fn)
	worker.Start()

	time.Sleep(35 * time.Millisecond)
	fmt.Println(worker.Err())

}

func TestCancel() {
	count := 9
	fn := func() error {
		fmt.Print(count, " ")
		count--
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	worker := NewWorker(fn)
	worker.Start()
	time.Sleep(105 * time.Millisecond)
	worker.Stop()

	fmt.Println()
	// 9 8 7 6 5 4 3 2 1 0
}

func TestAfterStop() {
	fn := func() error { return nil }

	worker := NewWorker(fn)
	worker.Start()
	worker.Start()
	worker.Start()
	worker.Start()
	worker.Start()
	time.Sleep(10 * time.Millisecond)

	fmt.Println(worker.Err())
}
