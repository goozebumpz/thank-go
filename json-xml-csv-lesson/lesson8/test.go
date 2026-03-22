package lesson8

import (
	"encoding/json"
	"fmt"
)

type Genre string

type Movie struct {
	Title  string
	Year   int
	Genres []Genre
}

func (g *Movie) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	type tag struct {
		Name string `json:"name"`
	}

	type Alias struct {
		Title  string `json:"name"`
		Year   int    `json:"released_at"`
		Genres []tag  `json:"tags"`
	}

	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	g.Year = a.Year
	g.Title = a.Title

	for _, tagInner := range a.Genres {
		g.Genres = append(g.Genres, Genre(tagInner.Name))
	}

	return nil
}

func Test() {
	const src = `{
        "name": "Interstellar",
        "released_at": 2014,
        "director": "Christopher Nolan",
        "tags": [
            { "name": "Adventure" },
            { "name": "Drama" },
            { "name": "Science Fiction" }
        ],
        "duration": "2h49m",
        "rating": "★★★★★"
    }`

	var movie Movie
	json.Unmarshal([]byte(src), &movie)
	fmt.Println(movie)
}
