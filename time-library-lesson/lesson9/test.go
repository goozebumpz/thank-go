package lesson9

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const test = `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

func ParseCell() {

}

func ParsePage(src string) []Task {
	var result []Task
	var date time.Time

	for _, str := range strings.Split(src, "\n") {
		words := strings.Split(str, " ")

		if len(words) == 0 {
			continue
		}

		if len(words) == 1 {
			newDate := words[0]
			parsedDate, err := time.Parse("02.01.2006", newDate)

			if err != nil {
				log.Fatal(err)
			} else {
				date = parsedDate
				continue
			}
		}

		if len(words) > 1 {
			title := strings.Join(words[3:], " ")
			from := words[0]
			to := words[2]
			parsedFrom := time.Time{}
			parsedTo := time.Time{}

			if len(from) == 4 {
				parsedFrom, _ = time.Parse("3:04", from)
			} else if len(from) == 5 {
				parsedFrom, _ = time.Parse("03:04", from)
			}

			if len(to) == 4 {
				parsedTo, _ = time.Parse("3:04", to)
			} else if len(to) == 5 {
				parsedTo, _ = time.Parse("03:04", to)
			}

			result = append(result, Task{Title: title, Date: date, Dur: parsedTo.Sub(parsedFrom)})
		}
	}

	return result
}

func Test() {
	res := ParsePage(test)
	fmt.Printf("%+v", res)
}
