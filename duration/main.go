package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	test := "1h30m"
	d, _ := time.ParseDuration(test)
	str := fmt.Sprintf("%s = %d min", test, int(math.Round(d.Minutes())))
	fmt.Println(str)
}
