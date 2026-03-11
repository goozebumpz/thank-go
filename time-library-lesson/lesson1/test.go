package lesson1

import (
	"fmt"
	"time"
)

func Test() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println("day", t.Day())
	fmt.Println("month", t.Month())
	fmt.Println("year", t.Year())
	fmt.Println("hour", t.Hour())
	fmt.Println("minute", t.Minute())
	fmt.Println("second", t.Second())
	fmt.Println("nanosecond", t.Nanosecond())

	fmt.Println(time.January == time.Month(1))
	fmt.Println(t.YearDay())
	hour, min, sec := t.Clock()
	fmt.Println(hour, min, sec)
}
