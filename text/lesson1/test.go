package lesson1

import (
	"fmt"
	"strings"
)

func Test() {
	test := "welcome to the secret shop"

	fmt.Println(strings.Contains(test, "dick"))
	fmt.Println(strings.HasPrefix(test, "welcome"))
	fmt.Println(strings.HasSuffix(test, "shop"))
	fmt.Println(strings.Index(test, "t"))
	fmt.Println(strings.LastIndex(test, "shop"))
	fmt.Println(strings.Count(test, "welcome"))
	fmt.Println(strings.Compare(test, "welcome"))

}
