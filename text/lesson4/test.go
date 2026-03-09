package lesson4

import (
	"fmt"
	"unicode"
)

func Test() {
	fmt.Println("is digit 0, ", unicode.IsDigit('0'))
	fmt.Println("is letter ы, ", unicode.IsLetter('ы'))
	fmt.Println("is letter <, ", unicode.IsLetter('<'))
	fmt.Println(unicode.Is(unicode.Latin, 'ы'))
}
