package lesson3

import (
	"fmt"
	"time"
)

func Test() {
	// 1
	utc := time.FixedZone("UTC", 0)
	t1 := time.Date(2022, 5, 24, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, 5, 24, 0, 0, 0, 0, utc)
	fmt.Println(t1.Equal(t2))

	// 2
	offsetSec := 3 * 3600
	utc3 := time.FixedZone("UTC+3", offsetSec)
	t1 = time.Date(2022, 5, 24, 0, 0, 0, 0, time.Local)
	t2 = time.Date(2022, 5, 24, 0, 0, 0, 0, utc3)
	fmt.Println(t1.Equal(t2))

	// 3
	paris, _ := time.LoadLocation("Europe/Paris")
	utc2 := time.FixedZone("UTC+2", 2*3600)
	t1 = time.Date(2022, 2, 22, 0, 0, 0, 0, paris)
	t2 = time.Date(2022, 2, 22, 0, 0, 0, 0, utc2)
	fmt.Println(t1.Equal(t2))

	// 4
	t := time.Date(2022, 2, 22, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	paris, _ = time.LoadLocation("Europe/Paris")
	fmt.Println(t.In(paris))

	ny, _ := time.LoadLocation("America/New_York")
	fmt.Println(t.In(ny))
}
