package main

import (
	"fmt"
	"time"
)

func testChanel() {
	messages := make(chan string)

	go func() {
		fmt.Println("B: Sending message...")
		messages <- "ping"
		fmt.Println("B: Message sent !")
	}()

	fmt.Println("A: Doing some work...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("A: Ready to receive a message...")

	<-messages
	fmt.Println("A: Messege received!")
	time.Sleep(100 * time.Millisecond)
}
