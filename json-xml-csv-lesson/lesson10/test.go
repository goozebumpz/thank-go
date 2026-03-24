package lesson10

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ex1() {
	peopleFile, err := os.Open("people.json")

	if err != nil {
		panic(err)
	}

	defer peopleFile.Close()
	dec := json.NewDecoder(bufio.NewReader(peopleFile))

	for {
		token, err := dec.Token()
		if err != nil {
			break
			panic(err)
		}

		fmt.Printf("%T: %v", token, token)

		if dec.More() {
			fmt.Print("...")
			fmt.Print("\n")
		}
	}
}

func Test() {
	ex1()
}
