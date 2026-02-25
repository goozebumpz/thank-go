package lesson3

import "errors"

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

type Queue struct {
	ch chan int
}

func (q Queue) Get(block bool) (int, error) {
	select {
	case v := <-q.ch:
		return v, nil
	default:
		if block {
			select {
			case v := <-q.ch:
				return v, nil
			}
		} else if !block {
			return 0, ErrEmpty
		}
		return 0, nil
	}
}

func (q Queue) Put(val int, block bool) error {
	select {
	case q.ch <- val:
		return nil
	default:
		if block {
			select {
			case q.ch <- val:
			}
			return nil
		} else if !block {
			return ErrFull
		}
		return nil
	}
}

func MakeQueue(n int) Queue {
	return Queue{make(chan int, n)}
}
