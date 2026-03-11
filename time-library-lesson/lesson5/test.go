package lesson5

import (
	"errors"
	"fmt"
	"time"
)

type TimeOfDay struct {
	date time.Time
}

func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	now := time.Now()

	return TimeOfDay{
		date: time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc),
	}
}

func (t TimeOfDay) Hour() int {
	return t.date.Hour()
}

func (t TimeOfDay) Minute() int {
	return t.date.Minute()
}

func (t TimeOfDay) Second() int {
	return t.date.Second()
}

func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %s", t.Hour(), t.Minute(), t.Second(), t.date.Location())
}

func (t TimeOfDay) Equal(other TimeOfDay) bool {
	l1 := t.date.Location().String()
	l2 := other.date.Location().String()

	if l1 != l2 {
		return false
	}

	return t.date.Equal(other.date)
}

func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.date.Location().String() != other.date.Location().String() {
		return false, errors.New("another locale")
	}

	return t.date.Before(other.date), nil
}

func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.date.Location().String() != other.date.Location().String() {
		return false, errors.New("another locale")
	}

	return t.date.After(other.date), nil
}

func Test() {
	hours := 1
	fmt.Println(fmt.Sprintf("%02d", hours))

}
