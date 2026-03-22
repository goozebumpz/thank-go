package lesson5

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	Weight    float64 `json:"-"`
	IsAwesome bool    `json:"is_awesome"`
}

func Test() {
	p := Person{"Roma", 26, 178.5, true}
	b, _ := json.MarshalIndent(p, "", "    ")

	fmt.Println(string(b))
}
