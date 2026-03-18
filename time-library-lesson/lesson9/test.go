package lesson9

import (
	"errors"
	"fmt"
	"sort"
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
	Date  time.Time
	Dur   time.Duration
	Title string
}

func parseDate(src string) (time.Time, error) {
	return time.Parse("02.01.2006", src)
}

func parseTasks(date time.Time, lines []string) ([]Task, error) {
	titles := map[string]time.Duration{}
	tasks := []Task{}

	for _, line := range lines {
		splt := strings.Split(line, " ")
		from, err := time.Parse("15:04", splt[0])

		if err != nil {
			return []Task{}, errors.New("error parse from")
		}

		to, err := time.Parse("15:04", splt[2])

		if err != nil {
			fmt.Printf("%v\n", err)
			return []Task{}, errors.New("error parse to")
		}

		if from.After(to) || from.Equal(to) {
			return []Task{}, errors.New("error invalidate period")
		}

		duration := to.Sub(from)
		desc := splt[3:]

		if len(desc) == 0 {
			return []Task{}, errors.New("error find description")
		}

		description := strings.Join(desc, " ")

		if durDesc, ok := titles[description]; ok {
			titles[description] = durDesc + duration
		} else {
			titles[description] = duration
		}
	}

	for title, dur := range titles {
		tasks = append(tasks, Task{Title: title, Dur: dur, Date: date})
	}

	return tasks, nil
}

func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur.Seconds() > tasks[j].Dur.Seconds()
	})
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

	sortTasks(tasks)

	return tasks, nil
}

func Test() {
	test := []string{"12:00 - 11:00 task"}

	res, err := parseTasks(time.Now(), test)
	fmt.Println(res, err)

}
