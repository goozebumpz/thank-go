package main

import (
	"fmt"
	"time"
)

func downloadFile(filename string) {
	fmt.Printf("Starting download %s \n", filename)
	time.Sleep(time.Second)
	fmt.Printf("Finished download %s \n", filename)
}

func test1() {
	startTime := time.Now()
	done := make(chan bool)

	go func() {
		downloadFile("file1.txt")
		done <- true
	}()
	go func() {
		downloadFile("file2.txt")
		done <- true

	}()
	go func() {
		downloadFile("file3.txt")
		done <- true
	}()

	for i := 0; i < 3; i++ {
		<-done
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("All downloads completed! Time elapsed: %s\n", elapsedTime)
}

func test2() {
	ch := make(chan string)

	go func() {
		ch <- "dick"
	}()

	msg := <-ch
	fmt.Println(msg)
}

func main() {
	test1()
}
