package lesson6

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatList(list []string) string {
	b := strings.Builder{}

	for i, str := range list {
		b.WriteString(strconv.Itoa(i+1) + ") ")
		b.WriteString(str)
		b.WriteString("\n")
	}

	return b.String()
}

func Test() {
	list := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
	}

	l := FormatList(list)

	fmt.Println(l)
}
