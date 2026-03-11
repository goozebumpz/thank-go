package lesson7

import (
	"fmt"
	"time"
)

func ex1() {
	// Чтобы прибавить ко времени продолжительность и получить новое время, используют метод Add():
	before := time.Date(2000, 2, 4, 0, 0, 0, 0, time.UTC)
	after := before.Add(7 * time.Hour)
	fmt.Println(after)

	after = after.Add(30 * time.Minute)
	fmt.Println(after)
	after = after.Add(45 * time.Second)
	fmt.Println(after)
	after = after.Add(20 * time.Nanosecond)
	fmt.Println(after)
	after = after.Add(-30 * time.Minute)
	fmt.Println(after)
}

func ex2() {
	// Метод AddDate() добавляет указанное количество лет, месяцев и дней:
	before := time.Date(2000, 2, 4, 0, 0, 0, 0, time.UTC)
	fmt.Println(before)
	before = before.AddDate(3, 2, 0)
	fmt.Println(before)
}

func ex3() {
	// Разность двух времен через метод Sub() возвращает продолжительность между ними:
	before := time.Date(2012, 4, 5, 0, 0, 0, 0, time.UTC)
	after := time.Date(2014, 4, 5, 0, 0, 0, 0, time.UTC)
	fmt.Println(after.Sub(before))

}

func Test() {
	ex1()
	ex2()
	ex3()
}
