package lesson7

import (
	"context"
	"fmt"
	"strings"
	"unicode"
)

type counter map[string]int
type pair struct {
	word  string
	count int
}

func countDigits(str string) int {
	count := 0

	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}

	return count
}

func countDigitsInWords(ctx context.Context, words []string) counter {
	select {
	case <-ctx.Done():
		return counter{}
	default:
		pending := submitWords(ctx, words)
		counted := countWords(ctx, pending)
		return fillStats(ctx, counted)
	}
}

func submitWords(ctx context.Context, words []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for _, word := range words {
			select {
			case <-ctx.Done():
				return
			case out <- word:
			}
		}
	}()

	return out
}

func countWords(ctx context.Context, in <-chan string) <-chan pair {
	out := make(chan pair)

	go func() {
		defer close(out)
		for word := range in {
			select {
			case <-ctx.Done():
				return
			case out <- pair{word: word, count: countDigits(word)}:
			}
		}
	}()

	return out
}

func fillStats(ctx context.Context, in <-chan pair) counter {
	stats := counter{}

	for p := range in {
		select {
		case <-ctx.Done():
			return counter{}
		default:
			stats[p.word] = p.count
		}
	}

	return stats
}

func Test() {
	phrase := "0ne 1wo thr33 4068"
	words := strings.Fields(phrase)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stats := countDigitsInWords(ctx, words)
	fmt.Println(stats)
}
