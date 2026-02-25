package main

import "fmt"

type Total struct {
	count  int
	amount int
}

func rangeGenNew(start, stop int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			out <- i
		}
	}()

	return out
}

func takeLucky(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			if num%7 == 0 || num%13 != 0 {
				out <- num
			}
		}
	}()

	return out
}

func mergeNext(chans ...<-chan int) <-chan int {
	out := make(chan int)
	done := make(chan bool)

	for _, ch := range chans {
		go func() {
			for ch != nil {
				select {
				case val, ok := <-ch:
					if ok {
						out <- val
					} else {
						ch = nil
						done <- true
					}
				}
			}
		}()
	}

	go func() {
		defer close(done)
		defer close(out)
		for range chans {
			<-done
		}
	}()

	return out
}

func sum(in <-chan int) <-chan Total {
	out := make(chan Total)

	go func() {
		defer close(out)
		total := Total{}

		for num := range in {
			total.amount += num
			total.count++
		}
		out <- total
	}()

	return out
}

func printTotal(in <-chan Total) {
	total := <-in
	fmt.Printf("Total of %d lucky numbers = %d\n", total.count, total.amount)
}

func test() {
	readerChan := rangeGen(1, 1000)
	luckyChans := make([]<-chan int, 4)
	for i := 0; i < 4; i++ {
		luckyChans[i] = takeLucky(readerChan)
	}
	mergedChan := mergeNext(luckyChans...)
	totalChan := sum(mergedChan)
	printTotal(totalChan)
}
