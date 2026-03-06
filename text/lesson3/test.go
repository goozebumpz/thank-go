package lesson3

import (
	"fmt"
	"strconv"
)

func Test() {
	n, err := strconv.Atoi("dick")

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(n)
}
