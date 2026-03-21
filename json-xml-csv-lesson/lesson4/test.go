package lesson4

import (
	"encoding/json"
	"fmt"
	"time"
)

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	result := ""
	instance := time.Duration(d)
	minutes := instance.Minutes()

	if minutes >= 60 {
		hours := int(minutes / 60)
		mins := int(minutes) - hours*60

		result += fmt.Sprintf("%dh", hours)

		if mins > 0 {
			result += fmt.Sprintf("%dm", mins)
		}
	} else {
		result = fmt.Sprintf("%dm", int(minutes))
	}

	return []byte(fmt.Sprintf(`"%s"`, result)), nil
}

type Rating int

func (r Rating) MarshalText() ([]byte, error) {
	res := ""

	for i := 0; i < 5; i++ {
		if i < int(r) {
			res += "★"
		} else {
			res += "☆"
		}
	}

	return []byte(fmt.Sprintf(`%s`, res)), nil
}

type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

func MarshalMovies(indent int, movies ...Movie) (string, error) {
	indentStr := ""

	if indent > 0 {
		for len(indentStr) < indent {
			indentStr += " "
		}

		b, err := json.MarshalIndent(movies, "", indentStr)

		if err != nil {
			return "", err
		}

		return string(b), nil
	}

	b, err := json.Marshal(movies)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func Test() {
	movie1 := Movie{"Interstellar", 2014, "Christopher Nolan", []string{"Adventure", "Drama", "Science", "Fiction"}, 10140000000000, 5}
	movie2 := Movie{
		Title:    "Pick",
		Year:     2022,
		Director: "Roman Repelling",
		Genres:   []string{"Garbage"},
		Duration: Duration(0),
		Rating:   2,
	}

	str, err := MarshalMovies(4, movie1, movie2)

	if err != nil {
		panic(err)
	}

	fmt.Println(str)
}

func TestSecond() {
	dur := Duration(60 * time.Minute)
	res, _ := json.Marshal(dur)
	fmt.Println(string(res))
}
