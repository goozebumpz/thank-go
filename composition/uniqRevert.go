package main

import (
	"fmt"
	"math/rand"
)

func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)

	go func() {
		defer fmt.Println("generate goroutine exited")
		defer close(out)
		for {
			select {
			case out <- randomWord(5):
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func takeUniq(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)

	checkUniq := func(str string) bool {
		m := make(map[rune]struct{}, len(str))

		for _, char := range str {
			if _, ok := m[char]; ok {
				return false
			}
			m[char] = struct{}{}
		}

		return true
	}

	go func() {
		defer fmt.Println("unique goroutine exited")
		defer close(out)

		for str := range in {
			if checkUniq(str) {
				select {
				case out <- str:
				case <-cancel:
					return
				}
			}
		}
	}()

	return out
}

type Original struct {
	from string
	to   string
}

func reverse(cancel <-chan struct{}, in <-chan string) <-chan Original {
	out := make(chan Original)

	reverseFunc := func(str string) Original {
		bytes := make([]byte, len(str))

		for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = str[j], str[i]
		}

		return Original{
			from: str,
			to:   string(bytes),
		}
	}

	go func() {
		defer fmt.Println("reverse goroutine exited")
		defer close(out)

		for str := range in {
			select {
			case out <- reverseFunc(str):
			case <-cancel:
				return
			}
		}
	}()

	return out
}

func mergeUniqRevert(cancel <-chan struct{}, c1, c2 <-chan Original) <-chan Original {
	out := make(chan Original)

	go func() {
		defer fmt.Println("merge goroutine exited")
		defer close(out)
		for c1 != nil || c2 != nil {
			select {
			case out <- <-c1:
			case out <- <-c2:
			case <-cancel:
				fmt.Println("merge get cancel")
				c1 = nil
				c2 = nil
				return
			}
		}

	}()

	return out
}

func printUniqRevert(cancel <-chan struct{}, in <-chan Original, count int) {
	defer fmt.Println("print revert goroutine exited")
	counter := 0

	for in != nil && counter < count {
		select {
		case out, ok := <-in:
			if ok {
				fmt.Printf("%s -> %s\n", out.from, out.to)
				counter++
			} else {
				in = nil
			}
		case <-cancel:
			return
		}
	}
}
