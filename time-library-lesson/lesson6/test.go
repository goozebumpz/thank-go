package lesson6

import (
	"fmt"
	"time"
)

func Test() {
	// 1
	//Продолжительность можно создать из строкового описания с помощью функции time.ParseDuration():
	d1, _ := time.ParseDuration("30s")
	fmt.Printf("%#v\n", d1)
	d2, _ := time.ParseDuration("2h36m22s")
	fmt.Printf("%#v\n", d2)

	//2
	// Продолжительность измеряется с точностью до наносекунд, но можно получить ее в микро-, милли- или обычных секундах:
	d3, _ := time.ParseDuration("567s")
	fmt.Printf("%#v\n", d3)
	fmt.Printf("Hours: %.5f \n", d3.Hours())
	fmt.Printf("Minutes: %.5f \n", d3.Minutes())
	fmt.Printf("Seconds: %.1f \n", d3.Seconds())
	fmt.Printf("Milliseconds: %d \n", d3.Milliseconds())
	fmt.Printf("Microseconds: %d \n", d3.Microseconds())
	fmt.Printf("Nanoseconds: %d \n", d3.Nanoseconds())

	// 3
	// Чтобы не парсить продолжительность, ее часто задают явно, с помощью готовых констант:
	d4 := 30 * time.Second
	fmt.Println(d4)

	d5 := 15 * time.Minute
	fmt.Println(d5)

	d6 := 2 * time.Hour
	fmt.Println(d6)

	//4
	// Продолжительности можно складывать и вычитать:
	d := d4 + d5 + d6
	fmt.Println(d)

	//5
	// А функция time.Since() возвращает продолжительность с указанного момента до текущего:
	before := time.Now()
	time.Sleep(time.Second)
	elapsed := time.Since(before)
	fmt.Println(elapsed)
}
