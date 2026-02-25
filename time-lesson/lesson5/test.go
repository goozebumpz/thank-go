package lesson5

import "time"

func After(dur time.Duration) <-chan time.Time {
	ch := make(chan time.Time, 1)

	go func() {
		time.Sleep(dur)
		ch <- time.Now()
		close(ch)
	}()

	return ch
}
