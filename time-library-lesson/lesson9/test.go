package lesson9

import (
	"errors"
	"strings"
	"time"
)

const TEXT = `15.04.2022
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
	Date     time.Time
	Duration time.Duration
	Title    string
}

func parseDate(src string) (time.Time, error) {
	return time.Parse("02.01.2006", src)
}

func parseTasks(date time.Time, src []string) ([]Task, error) {
	titles := map[string]bool{}

	for _, line := range src {
		description := strings.Split(line, " ")

	}

	return []Task{}, nil
}

func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])

	if err != nil {
		return []Task{}, errors.New("error parsing date")
	}

	tasks, err := parseTasks(date, lines[1:])

	if err != nil {
		return []Task{}, errors.New("error parsing tasks")
	}

	return tasks, nil
}

func Test() {
	ParsePage(TEXT)
}
