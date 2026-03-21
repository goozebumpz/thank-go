package lesson1

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	Weight    float64
	IsAwesome bool
	secret    string
}

func ex1() {
	Ann := Person{
		Name:      "Ann",
		Age:       26,
		Weight:    50,
		IsAwesome: true,
		secret:    "My sister",
	}

	b, err := json.Marshal(Ann)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func ex2() {
	Ann := Person{
		Name:      "Ann",
		Age:       26,
		Weight:    50,
		IsAwesome: true,
		secret:    "My sister",
	}

	b, err := json.MarshalIndent(Ann, "", "    ")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func Test() {
	ex2()
}
