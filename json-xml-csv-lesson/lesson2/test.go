package lesson2

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Birthdate time.Time
	D         Dick
}

func ex1() {
	date, _ := time.Parse("2006-01-02", "2000-04-02")
	ann := Person{
		Name:      "Ann",
		Birthdate: date,
		D: Dick{
			sm: 20,
		},
	}

	b, err := json.Marshal(ann)
	fmt.Println(err, string(b))
}

type Dick struct {
	sm int
}

func (d *Dick) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

func ex2() {
	test := []string{"niga", "well", "done"}
	b, _ := json.Marshal(test)
	fmt.Println(string(b))
}

func ex3() {
	m := map[string]int{
		"roma": 21,
		"dima": 30,
		"edya": 35,
	}

	b, _ := json.Marshal(m)
	fmt.Println(string(b))

}

func Test() {
	ex1()
	ex2()
	ex3()
}
