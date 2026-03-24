package lesson9

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string
}

func ex1() {
	f, err := os.Open("people.jl")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)
	for {
		var person Person
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

func Test() {
	ex1()
}
