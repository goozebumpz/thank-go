package lesson7

import "time"

func Delay(dur time.Duration, fn func()) func() {
	cancelCh := make(chan struct{}, 1)
	isCanceled := false

	go func() {
		timer := time.NewTimer(dur)
		select {
		case <-timer.C:
			fn()
		case <-cancelCh:

			return
		}
	}()

	return func() {
		if !isCanceled {
			cancelCh <- struct{}{}
			close(cancelCh)
			isCanceled = true
		}
	}
}
