package lesson6

import (
	"encoding/json"
	"fmt"
)

const str = `
{
    "name": "Alice",
    "is_awesome": true,
    "residence": {
        "country": "France",
        "city": "Paris"
    },
    "friends": [
        { "name": "Emma" },
        { "name": "Grace" }
    ]
}`

func ex1() {
	fmt.Println(json.Valid([]byte(str)))
	// true
}

type Address struct {
	Country string
	City    string
}

type Person struct {
	Name      string
	IsAwesome bool `json:"is_awesome"`
	Residence Address
	Friends   []*Person
}

func ex2() {
	var alice Person
	err := json.Unmarshal([]byte(str), &alice)

	fmt.Println(err, alice)
}

type AncientNumber string

func (an *AncientNumber) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var n int
	err := json.Unmarshal(data, &n)
	if err != nil {
		return err
	}

	switch {
	case n <= 0:
		*an = "impossible!"
	case n == 1:
		*an = "one"
	case n == 2:
		*an = "two"
	case n > 2:
		*an = "many"
	}
	return nil
}

func ex3() {
	var n AncientNumber

	err := json.Unmarshal([]byte("1"), &n)
	fmt.Println(err, n)
	// <nil> one

	err = json.Unmarshal([]byte("2"), &n)
	fmt.Println(err, n)
	// <nil> two

	err = json.Unmarshal([]byte("42"), &n)
	fmt.Println(err, n)
	// <nil> many

	err = json.Unmarshal([]byte("-1"), &n)
	fmt.Println(err, n)
}

func Test() {
	ex1()
	ex2()
	ex3()
}
