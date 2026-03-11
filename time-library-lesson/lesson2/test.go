package lesson2

import (
	"fmt"
	"time"
)

func isLeapYear(year int) bool {
	t := time.Date(year, 12, 31, 0, 0, 0, 0, time.Local)
	return t.YearDay() == 366
}

func Test() {
	fmt.Println(isLeapYear(2020))
	fmt.Println(isLeapYear(2022))

}
