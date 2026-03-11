package lesson4

import (
	"fmt"
	"time"
)

func Test() {

	// 1
	t1 := time.Date(2022, 5, 24, 17, 45, 22, 0, time.UTC)
	t2 := time.Date(2022, 5, 24, 20, 45, 22, 0, time.FixedZone("UTC+3", 3*60*60))

	fmt.Println(t1 == t2)
	fmt.Println(t1.Equal(t2))

	// 2
	t3 := time.Date(2022, 5, 25, 0, 0, 0, 0, time.UTC)
	t4 := time.Date(2022, 5, 25, 17, 45, 22, 0, time.UTC)
	fmt.Println(t3.After(t4))
	fmt.Println(t3.Before(t4))

	// 3
	var timeZero time.Time
	fmt.Println(timeZero)
	fmt.Println(timeZero.IsZero())
}
