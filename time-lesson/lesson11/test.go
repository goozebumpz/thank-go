package lesson11

import (
	"errors"
	"time"
)

var ErrCanceled error = errors.New("canceled")

func Throttle(limit int, fn func()) (handle func() error, cancel func()) {
	isCanceled := make(chan struct{})
	jobs := make(chan func(), limit)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-isCanceled:
				return
			case <-ticker.C:
			default:
			}
		}
	}()

	handle = func() error {
		select {
		case <-isCanceled:
			return ErrCanceled
		default:
			jobs <- fn
			return nil
		}
	}

	cancel = func() {
		select {
		case <-isCanceled:
			return
		default:
			close(isCanceled)
		}
	}

	return handle, cancel
}

func Throttle2(limit int, fn func()) (handle func() error, cancel func()) {
	canceled := make(chan struct{})
	worked := make(chan struct{}, 1)

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(limit))
		defer ticker.Stop()

		for {
			select {
			case <-canceled:
				return
			case <-ticker.C:
				select {
				case <-canceled:
					return
				default:
					worked <- struct{}{}
				}
			}
		}
	}()

	handle = func() error {
		select {
		case <-canceled:
			return ErrCanceled
		case <-worked:
			select {
			case <-canceled:
				return ErrCanceled
			default:
				go fn()
			}
		}
		return nil
	}

	cancel = func() {
		select {
		case <-canceled:
			return
		default:
			close(canceled)
		}
	}

	return handle, cancel
}

// Grok shit
//func Throttle(limit int, fn func()) (handle func() error, cancel func()) {
//
//	isCanceled := make(chan struct{})
//	sem := make(chan struct{}, limit) // семафор на rate одновременно
//
//	go func() {
//		ticker := time.NewTicker(time.Second)
//		defer ticker.Stop()
//
//		for {
//			select {
//			case <-isCanceled:
//				return
//			case <-ticker.C:
//				for i := 0; i < limit; i++ {
//					select {
//					case <-isCanceled:
//						return
//					case <-sem:
//					default:
//						// больше нечего освобождать
//						break
//					}
//				}
//			}
//		}
//	}()
//
//	handle = func() error {
//		select {
//		case <-isCanceled:
//			return ErrCanceled
//		case sem <- struct{}{}: // ждём свободный слот
//			fn()
//			return nil
//		}
//	}
//
//	cancel = func() {
//		select {
//		case <-isCanceled:
//			return
//		default:
//			close(isCanceled)
//		}
//	}
//
//	return handle, cancel
//}
