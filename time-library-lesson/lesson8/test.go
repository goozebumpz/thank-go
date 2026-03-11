package lesson8

import (
	"fmt"
	"time"
)

func ex1() {
	layout := "02.01.2006"
	value := "12.01.2020"

	time, err := time.Parse(layout, value)

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("Result: %v", time)
	}
}

func ex2() {
	value := "24.05.2022 17:45"
	time, _ := time.Parse("02.01.2006 15:04", value)
	fmt.Println(time)
}

func ex3() {
	s := "24.05.2022 17:45:22.951205+03:00"
	t, _ := time.Parse("02.01.2006 15:04:05.999999-07:00", s)
	fmt.Println(t)
}

func ex4() {
	s := "2022-05-24 17:45:22"
	t, _ := time.Parse(time.DateTime, s)
	fmt.Println(t)

	s = "2022-04-02"
	t, _ = time.Parse(time.DateTime, s)
	fmt.Println(s)

	s = "14:25:00"
	t, _ = time.Parse(time.TimeOnly, s)
	fmt.Println(t)
}

func Test() {
	ex1()
	ex2()
	ex3()
	ex4()
}
