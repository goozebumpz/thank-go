package main

import (
	"fmt"
	"text-lesson/lesson12"
)

func main() {
	res := lesson12.Slugify("We haven’t killed 90% of all plankton")
	fmt.Println(res)
}
