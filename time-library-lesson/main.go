package main

import (
	"fmt"
	"time"
	"time-library-lesson/lesson9"
)

func main() {
	start := time.Now()
	after := time.Now()
	fmt.Println(start.Sub(after))
	fmt.Println(time.Parse("03:04", "04:12"))
	lesson9.Test()
}
