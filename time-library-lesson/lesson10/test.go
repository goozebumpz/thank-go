package lesson10

import (
	"fmt"
	"time"
)

func ex1() {
	// Go преобразует обычное время в юникс-время с любой точностью — от секунд до наносекунд:
	t, _ := time.Parse(time.RFC3339, "2022-05-24T17:45:22.951205+03:00")
	fmt.Println("Unix: ", t.Unix())
	fmt.Println("UnixMilli: ", t.UnixMilli())
	fmt.Println("UnixMicro: ", t.UnixMicro())
	fmt.Println("UnixNano: ", t.UnixNano())
}

func ex2() {
	// И обратно — из юникс-времени в обычное:
	// из секунд
	t := time.Unix(1653403522, 0)
	fmt.Println(t)
	// из наносекунд
	t = time.Unix(0, 1653403522951205123)
	fmt.Println(t)
	// из секунд и наносекундного остатка
	t = time.Unix(1653403522, 951205123)
	fmt.Println(t)
	// из миллисекунд
	t = time.UnixMilli(1653403522951)
	fmt.Println(t)
	// из микросекунд
	t = time.UnixMicro(1653403522951205)
	fmt.Println(t)

}

func Test() {
	ex1()
	ex2()
}
