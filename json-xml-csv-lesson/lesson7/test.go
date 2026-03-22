package lesson7

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
}

const src = `{
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

	var alice Person
	err := json.Unmarshal([]byte(src), &alice)
	fmt.Println(err, alice)
}

func ex2() {
	// decode to map
	var alice any
	err := json.Unmarshal([]byte(src), &alice)
	fmt.Printf("%v %T\n", err, alice)
	fmt.Println(alice)
}

func ex3() {
	var alice any
	json.Unmarshal([]byte(src), &alice)
	m := alice.(map[string]any)
	name := m["name"].(string)
	fmt.Printf("Name: %s", name)

	if m["is_awesome"].(bool) {
		fmt.Println("Alice is awesome")
	} else {
		fmt.Println("Alice is okay")
	}

	friends := m["friends"].([]any)
	for idx, friend := range friends {
		m := friend.(map[string]any)
		name := m["name"].(string)
		fmt.Printf("- friend #%d: %s\n", idx+1, name)
	}
}

func Test() {
	ex1()
	ex2()
	ex3()
}
