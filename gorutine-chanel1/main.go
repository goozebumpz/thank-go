package main

import (
	"fmt"
	"strings"
	"time"
)

// encode кодирует строку шифром Цезаря
func encode(str string) string {
	// начало решения

	submitter := func(str string) <-chan string {
		ch := make(chan string)
		go func() {
			words := strings.Fields(str)
			for _, word := range words {
				ch <- word
			}
			close(ch)
		}()
		return ch
	}

	encoder := func(ch1 <-chan string) <-chan string {
		ch2 := make(chan string)
		go func() {
			for word := range ch1 {
				ch2 <- encodeWord(word)
			}
			close(ch2)
		}()
		return ch2
	}

	receiver := func(ch <-chan string) []string {
		words := []string{}
		for word := range ch {
			words = append(words, word)
		}
		return words
	}

	// конец решения

	pending := submitter(str)
	encoded := encoder(pending)
	words := receiver(encoded)
	return strings.Join(words, " ")
}

// encodeWord кодирует слово шифром Цезаря
func encodeWord(word string) string {
	const shift = 13
	const char_a byte = 'a'
	encoded := make([]byte, len(word))
	for idx, char := range []byte(word) {
		delta := (char - char_a + shift) % 26
		encoded[idx] = char_a + delta
	}
	return string(encoded)
}

func test() {
	src := "go is awesome"
	res := encode(src)
	fmt.Println(res)
}

func test1() {
	ch := make(chan int)

	go func() {
		for val1 := range ch {
			fmt.Println("Gor 1", val1)
		}
	}()

	go func() {
		for val2 := range ch {
			fmt.Println("Gor 2", val2)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	time.Sleep(3 * time.Second)
}

func main() {
	test1()
}
