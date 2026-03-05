package lesson12

import "sync/atomic"

type Total struct {
	atomic.Int32
}

func (t *Total) Increment() {
	t.Add(1)
}

func (t *Total) Value() int {
	return int(t.Load())
}
