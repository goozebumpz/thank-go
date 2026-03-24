package lesson11

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func ex1() {
	f, err := os.Open("people.json")

	if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(bufio.NewReader(f))

	if _, err := dec.Token(); err != nil {
		panic(err)
	}

	var person Person

	for dec.More() {
		err := dec.Decode(&person)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(person)
	}
}

func ex2() {
	people := []Person{
		{Name: "Alice", Age: 24},
		{Name: "Grace", Age: 22},
		{Name: "Emma", Age: 29},
	}

	f, err := os.Create("people.jl")

	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	enc := json.NewEncoder(w)

	for _, person := range people {
		err := enc.Encode(person)
		if err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}
}

func Test() {
	ex1()
	ex2()
}
