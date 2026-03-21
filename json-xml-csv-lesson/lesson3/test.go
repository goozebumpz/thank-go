package lesson3

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Country string
	City    string
}

type Person struct {
	Name      string
	Residence Address
}

func ex1() {
	address := Address{"France", "Paris"}
	alice := Person{"Alice", address}
	b, _ := json.MarshalIndent(alice, "", "    ")
	fmt.Println(string(b))
}

type PersonPointerAddress struct {
	Name      string
	Residence *Address
}

func ex2() {
	address := Address{"Russia", "Moscow"}
	roma := PersonPointerAddress{"Roma", &address}
	b, _ := json.Marshal(roma)
	fmt.Println(string(b))
}

func ex3() {
	denis := Person{
		Name: "Denis",
	}
	p, _ := json.Marshal(denis)
	fmt.Println(string(p))
}

func ex4() {
	dasha := PersonPointerAddress{
		Name: "Dasha",
	}
	p, _ := json.MarshalIndent(dasha, "", "    ")
	fmt.Println(string(p))
}

type UserWithFriends struct {
	Name    string
	Friends []*Person
}

func ex5() {
	dmitriy := Person{
		Name: "Dmitry",
	}
	andrey := Person{
		Name: "Andrey",
	}
	roma := UserWithFriends{
		Name:    "Roma",
		Friends: []*Person{&dmitriy, &andrey},
	}

	b, _ := json.MarshalIndent(roma, "", "    ")

	fmt.Println(string(b))
}

func Test() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}
