// Package words
package main

// не удаляйте импорты, они используются при проверке
import (
	"fmt"
	"strings"
)

// Words работает со словами в строке.
type Words map[string]int

// MakeWords создает новый экземпляр Words.
func MakeWords(s string) Words {
	words := strings.Fields(s)
	size := len(words)
	mapWords := make(map[string]int, size)

	for i, word := range words {
		if _, has := mapWords[word]; has {
			continue
		}

		mapWords[word] = i
	}

	return mapWords
}

// Index возвращает индекс первого вхождения слова в строке,
// или -1, если слово не найдено.
func (w Words) Index(word string) int {
	index, has := w[word]

	if !has {
		return -1
	}

	return index
}

func main() {
	words := MakeWords("in a coat of gold or a coat of red")
	fmt.Println(words)
	fmt.Println(words.Index("a"))
}
