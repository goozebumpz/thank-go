package lesson10

import (
	"time"
)

func Schedule(dur time.Duration, fn func()) func() {
	canceled := make(chan struct{})
	ticker := time.NewTicker(dur)

	go func() {

		select {
		case <-canceled:
			return
		default:
			for {
				select {
				case <-canceled:
					return
				case <-ticker.C:
					select {
					case <-canceled:
						return
					default:
						select {
						case <-canceled:
							return
						default:
							fn()
						}
					}
				}
			}
		}
	}()

	return func() {
		select {
		case <-canceled:
			return
		default:
			ticker.Stop()
			close(canceled)
		}
	}
}
