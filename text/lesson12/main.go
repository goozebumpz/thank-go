package lesson12

import (
	"strings"
)

func Slugify(src string) string {
	b := strings.Builder{}
	sliceBytes := []byte(src)

	for index, char := range sliceBytes {
		if char > 127 {
			continue
		}

		isDigit := char >= '0' && char <= '9'
		isLowerChar := char >= 'a' && char <= 'z'
		isUpperChar := char >= 'A' && char <= 'Z'
		isHyphen := char == '-'

		if isDigit || isLowerChar || isHyphen {
			b.WriteByte(char)
			continue
		}

		if isUpperChar {
			b.WriteByte(char + 32)
			continue
		}

		if char == 32 || char == 45 || char == 46 {
			b.WriteByte(45)
			continue
		}

		if index < len(sliceBytes)-1 {
			b.WriteByte(45)
		}

	}

	return b.String()
}

func Slugify2(src string) string {
	b := strings.Builder{}
	sliceBytes := []byte(src)
	result := make([]string, 0)

	for _, char := range sliceBytes {
		if char > 127 {
			b.WriteByte(45)
			continue
		}

		if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || char == 45 {
			b.WriteByte(char)
			continue
		}

		if char >= 'A' && char <= 'Z' {
			b.WriteByte(char + 32)
			continue
		}

		if b.Len() > 0 {
			result = append(result, b.String())
			b.Reset()
		}
	}

	if b.Len() > 0 {
		result = append(result, b.String())
		b.Reset()
	}

	return strings.Join(result, "-")
}

func Slugify3(src string) string {
	b := strings.Builder{}
	sliceBytes := []byte(src)
	lastWasSafe := false

	for _, char := range sliceBytes {
		if char > 127 {
			continue
		}

		isDigit := char >= '0' && char <= '9'
		isLower := char >= 'a' && char <= 'z'
		isUpper := char >= 'A' && char <= 'Z'
		isHyphen := char == '-'

		if isDigit || isLower || isHyphen {
			b.WriteByte(char)
			lastWasSafe = true
			continue
		}

		if isUpper {
			b.WriteByte(char + 32)
			lastWasSafe = true
			continue
		}

		if lastWasSafe && (char == ' ' || char == '.' || char == '_') {
			b.WriteByte('-')
			lastWasSafe = false
		}
	}

	result := b.String()
	if len(result) > 0 && result[len(result)-1] == '-' {
		result = result[:len(result)-1]
	}

	return result
}
