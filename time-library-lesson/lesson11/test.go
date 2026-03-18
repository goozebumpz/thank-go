package lesson11

import (
	"fmt"
	"time"
)

func ex1() {
	// m=+... — это монотонное время (monotonic time).
	// Монотонное время — это локальное время процесса в секундах.
	// Сразу после старта процесса оно равно 0, а дальше увеличивается на 0.000000001 с каждой прошедшей наносекундой:
	t := time.Now()
	fmt.Println(t)
}

func ex2() {
	fmt.Println(time.Now())
	time.Sleep(50 * time.Millisecond)
	fmt.Println(time.Now())
	time.Sleep(50 * time.Millisecond)
	fmt.Println(time.Now())
}

func ex3() {
	now := time.Now()
	time.Sleep(50 * time.Millisecond)
	elapsed := time.Since(now)
	fmt.Println(elapsed)
}

func ex4() {
	deadline := time.Now().Add(60 * time.Second)
	time.Sleep(time.Second)
	fmt.Println(time.Until(deadline))
}

func ex5() {
	// Монотонная часть существует только у времени, полученного через time.Now().
	//Если создать время через time.Date() или time.Parse() — ее не будет:
	t1 := time.Date(2022, 5, 24, 17, 45, 22, 951205000, time.Local)
	fmt.Println(t1)
	t2, _ := time.Parse(time.RFC3339, "2022-05-24T17:45:22.951205+03:00")
	fmt.Println(t2)
}

func ex6() {
	// Чтобы «отрезать» монотонную часть, используют метод Round():
	t := time.Now()
	fmt.Println(t)

	wall := t.Round(0)
	fmt.Println(wall)
}

func Test() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}
