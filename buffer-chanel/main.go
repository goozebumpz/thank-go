package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func test() {
	stream := make(chan bool, 1)

	send := func() {
		fmt.Println("Sender ready to send")
		stream <- true
		fmt.Println(len(stream), 2)
		fmt.Println("Send !")
	}

	receive := func() {
		fmt.Println("Receiver not ready yet")
		fmt.Println(len(stream), 1)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Receiver not ready to receive...")

		<-stream
		fmt.Println("Received!")
	}

	var wg sync.WaitGroup

	wg.Go(send)
	wg.Go(receive)
	wg.Wait()

}

func await(f func() any) any {
	ch := make(chan any, 1)

	go func() {
		ch <- f()
	}()

	return <-ch
}

func gather(funcs []func() any) []any {
	operationCh := make(chan struct {
		index  int
		result any
	}, len(funcs))
	done := make(chan []any)

	for i, f := range funcs {
		go func() {
			operationCh <- struct {
				index  int
				result any
			}{i, f()}
		}()
	}

	go func() {
		result := make([]any, len(funcs))
		counter := 0

		for r := range operationCh {
			result[r.index] = r.result
			counter += 1

			if counter == len(funcs) {
				close(operationCh)
			}
		}

		done <- result
	}()

	return <-done
}

func test2() {
	slowpoke := func() any {
		fmt.Println("I'm so...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("pock")
		return "okey"
	}

	result := await(slowpoke)

	fmt.Println(result.(string))
}

func testGather() {
	func1 := func() any {
		return 1
	}
	func2 := func() any {
		return "dick"
	}
	func3 := func() any {
		return "pick"
	}

	funcs := []func() any{func1, func2, func3}
	res := gather(funcs)
	fmt.Println(res)
}

func say(pool chan<- int, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s \n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}

	pool <- id
}

func test3() {
	phrases := []string{
		"welocme to the secret shop", "military to step boundaries", "waffen is good deal", "matter is begone", "Welow woki snow is poki",
	}

	pool := make(chan int, 2)
	pool <- 1
	pool <- 2

	for _, phrase := range phrases {
		id := <-pool
		go say(pool, id, phrase)
	}

	<-pool
}

func say2(done chan<- struct{}, pending <-chan string, id int) {
	for phrase := range pending {
		for _, word := range strings.Fields(phrase) {
			fmt.Printf("Worker #%d says: %s...\n", id, word)
			dur := time.Duration(rand.Intn(100)) * time.Millisecond
			time.Sleep(dur)
		}
	}

	done <- struct{}{}
}

func test4() {
	phrases := []string{
		"welocme to the secret shop", "military to step boundaries", "waffen is good deal",
	}

	pending := make(chan string)

	go func() {
		for _, phrase := range phrases {
			pending <- phrase
		}
		close(pending)
	}()

	done := make(chan struct{})

	go say2(done, pending, 1)
	go say2(done, pending, 2)

	<-done
}

func makePool(n int, handler func(int, string)) (func(string), func()) {
	pool := make(chan int, n)

	for i := 0; i < n; i++ {
		pool <- i
	}

	handle := func(s string) {
		id := <-pool

		go func() {
			handler(id, s)
			pool <- 1
		}()
	}

	wait := func() {
		for i := 0; i < n; i++ {
			<-pool
		}
	}

	return handle, wait

}

func say3(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s \n", id, word)
	}
}

func main() {
	phrases := []string{
		"welcome to the secret shop", "military to step boundaries", "waffen ya is good deal",
	}

	handle, wait := makePool(2, say3)

	for _, phrase := range phrases {
		handle(phrase)
	}

	wait()
}
