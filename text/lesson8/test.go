package lesson8

import (
	"fmt"
	"regexp"
)

func Test() {
	re := regexp.MustCompile(`\d+`)
	s := "2050-11-05 is November 5th, 2050"

	// подходит ли строка под шаблон?
	ok := re.MatchString(s)
	fmt.Println(ok)

	// первое совпадение с шаблоном
	first := re.FindString(s)
	fmt.Println(first)

	// индекс начала и окончания
	// первого совпадения
	idx := re.FindStringIndex(s)
	fmt.Println(idx)

	// N совпадений с шаблоном (в данном случае 3)
	// если указать -1 - вернет все совпадения
	three := re.FindAllString(s, 3)
	fmt.Println(three)

	// индексы совпадений
	indices := re.FindAllStringIndex(s, 3)
	fmt.Println(indices)
}

func TestGroup() {
	re := regexp.MustCompile(`(\d\d\d\d)-(\d\d)-(\d\d)`)
	s := "2050-11-05 is November 5th, 2050"

	match := re.FindString(s)
	fmt.Println(match)

	groups := re.FindStringSubmatch(s)
	fmt.Println(groups)

	indices := re.FindStringSubmatchIndex(s)
	fmt.Println(indices)

}
