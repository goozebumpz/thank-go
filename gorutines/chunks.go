package main

import (
	"fmt"
	"strings"
	"time"
)

func testChunks() {
	str := "one,two,free,,four"
	in := make(chan string)

	go func() {
		words := strings.Split(str, ",")
		for _, word := range words {
			in <- word
		}
	}()

	for {
		word := <-in
		if word != "" {
			fmt.Printf("%s ", word)
		}
	}
}

func testChunks2() {
	str := "one,two,,four"
	in := make(chan string)

	go func() {
		words := strings.Split(str, ",")

		for _, word := range words {
			in <- word
		}
		close(in)
	}()

	for {
		word, ok := <-in

		if !ok {
			break
		}

		if word != "" {
			fmt.Printf("%s ", word)
		}
	}
}

func testRangeChanel() {
	ch := make(chan string)

	go func() {
		words := strings.Split("Welcome to the secret shop", " ")
		for _, word := range words {
			time.Sleep(1 * time.Second)
			ch <- word
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
		fmt.Println("я же негр я же работаю")
	}
}
