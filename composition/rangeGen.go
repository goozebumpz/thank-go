package main

import (
	"fmt"
	"sync"
	"time"
)

func rangeGen(start, stop int) <-chan int {
	out := make(chan int)

	go func() {
		for i := start; i < stop; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func rangeGen2(cancel <-chan struct{}, start, stop int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			select {
			case out <- i:
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func rangeGen3(start, stop int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			time.Sleep(50 * time.Millisecond)
			out <- i
		}
	}()

	return out
}

func testRangeGen1() {
	generated := rangeGen(10, 20)

	for num := range generated {
		fmt.Println(num)
	}
}

func testRangeGen2() {
	generated := rangeGen(10, 20)

	for num := range generated {
		fmt.Println(num)

		if num == 15 {
			break
		}
	}
}

func testRangeGen3() {
	cancel := make(chan struct{})
	defer close(cancel)

	generated := rangeGen2(cancel, 0, 10)

	for val := range generated {
		fmt.Println(val)
		if val == 15 {
			break
		}
	}
}

func merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)

	func() {
		defer close(out)
		for val := range in1 {
			out <- val
		}

		for val := range in2 {
			out <- val
		}
	}()

	return out
}

func testRangeGen4() {
	start := time.Now()
	in1 := rangeGen3(1, 10)
	in2 := rangeGen3(20, 30)

	merged := merge(in1, in2)

	for val := range merged {
		fmt.Print(val, " ")
	}

	fmt.Println("Took 1", time.Since(start))
}

func merge2(in1, in2 <-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Go(func() {
		for val := range in1 {
			out <- val
		}
	})

	wg.Go(func() {
		for val := range in2 {
			out <- val
		}
	})

	go func() {
		wg.Wait()
		fmt.Println("is wait")
		close(out)
	}()

	return out
}

func testRangeGen5() {
	in1 := rangeGen3(1, 10)
	in2 := rangeGen3(20, 30)

	merged := merge2(in1, in2)

	for val := range merged {
		fmt.Print(val, " ")
	}

	fmt.Println("Dick")
}

func merge3(in1, in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for in1 != nil || in2 != nil {
			select {
			case val1, ok := <-in1:
				if ok {
					out <- val1
				} else {
					in1 = nil
				}

			case val2, ok := <-in2:
				if ok {
					out <- val2
				} else {
					in2 = nil
				}
			}
		}
	}()

	return out
}

//func merge(in1, in2 <-chan int) <-chan int {
//    out := make(chan int)
//    go func() {
//        defer close(out)
//        for {
//            select {
//            case out <- <-in1:
//            case out <- <-in2:
//            }
//        }
//    }()
//    return out
//}

func testRangeGen6() {
	in1 := rangeGen3(1, 10)
	in2 := rangeGen3(20, 30)

	merged := merge3(in1, in2)

	for val := range merged {
		fmt.Print(val, " ")
	}
}

func testMyMerge() {
	in1 := rangeGen3(1, 10)
	in2 := rangeGen3(20, 30)
	in3 := rangeGen3(40, 60)
	in4 := rangeGen3(70, 100)
	in5 := rangeGen3(1200, 1201)

	merged := myMerge(in1, in2, in3, in4, in5)

	for val := range merged {
		fmt.Print(val, " ")
	}
}

func myMerge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	done := make(chan int)

	for i := 0; i < len(channels); i++ {
		go func() {
			ch := channels[i]
			for ch != nil {
				select {
				case val, ok := <-ch:
					if ok {
						out <- val
					} else {
						done <- 1
						ch = nil
					}
				}
			}
		}()
	}

	go func() {
		counter := 0
		for counter < len(channels) {
			<-done
			counter++
		}
		close(out)
		close(done)
	}()

	return out
}
