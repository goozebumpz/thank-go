package lesson8

import (
	"fmt"
	"sync"
)

type Counter struct {
	m    map[string]int
	lock *sync.Mutex
}

func (c *Counter) Increment(str string) {
	c.lock.Lock()
	c.m[str]++
	c.lock.Unlock()
}

func (c *Counter) Value(str string) int {
	c.lock.Lock()
	defer c.lock.Unlock()

	if value, has := c.m[str]; has {
		return value
	}
	return 0
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.Lock()
	for key, value := range c.m {
		fn(key, value)
	}
	c.lock.Unlock()
}

func NewCounter() *Counter {
	return &Counter{
		m:    map[string]int{},
		lock: &sync.Mutex{},
	}
}

func Test() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
