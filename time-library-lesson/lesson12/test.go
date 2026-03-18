package lesson12

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func asLegacyDate(t time.Time) string {
	seconds := t.Unix()
	nanoseconds := t.UnixNano()
	strSeconds := strconv.Itoa(int(seconds))
	strNanoseconds := strconv.Itoa(int(nanoseconds))
	strNanoseconds = strings.TrimPrefix(strNanoseconds, strSeconds)
	strNanoseconds = strings.TrimRight(strNanoseconds, "0")
	hasSuffix := strings.HasSuffix(strNanoseconds, ".")

	if hasSuffix || strNanoseconds == "" {
		strNanoseconds = "0"
	}

	return strSeconds + "." + strNanoseconds
}

func parseLegacyDate(d string) (time.Time, error) {
	if !strings.Contains(d, ".") {
		return time.Time{}, errors.New("format not support")
	}

	slice := strings.Split(d, ".")
	seconds, err := strconv.Atoi(slice[0])

	if err != nil {
		return time.Time{}, err
	}

	strNanoseconds := slice[1]

	if strNanoseconds == "" {
		return time.Time{}, errors.New("error empty nanoseconds")
	}

	normalizeNanoseconds := func(str string) (int64, error) {
		if len(str) == 1 {
			return strconv.ParseInt(str, 10, 64)
		}

		if strings.HasPrefix(str, "0") {
			str = strings.TrimLeft(str, "0")

			return strconv.ParseInt(str, 10, 64)
		}

		for len(str) < 9 {
			str = str + "0"
		}

		return strconv.ParseInt(str, 10, 64)
	}

	nanoseconds, err := normalizeNanoseconds(strNanoseconds)

	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(int64(seconds), nanoseconds).UTC(), nil
}

func Test() {
	res1, err := parseLegacyDate("3600.000000001")
	fmt.Println(err)
	fmt.Println(res1)

}
