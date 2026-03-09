package lesson10

import (
	"os"
	"text/template"
)

func Test() {
	text := "Alise: - {{.}}\n"
	tpl := template.New("value")
	tpl = template.Must(tpl.Parse(text))

	tpl.Execute(os.Stdout, "Привет!")
	tpl.Execute(os.Stdout, "Как дела?")
	tpl.Execute(os.Stdout, "Пока!")
}

func Test1() {
	text := `Сейчас {{.Time}}, {{.Day}} 
	{{if .Sunny -}} Солнечно! {{- else -}} Пасмурно :-/ {{- end}}
`

	tmpl := template.New("greeting")
	tmpl = template.Must(tmpl.Parse(text))

	type State struct {
		Time  string
		Day   string
		Sunny bool
	}

	state := State{"21:00", "Четверг", true}
	tmpl.Execute(os.Stdout, state)

	state = State{"21:00", "Пятница", false}
	tmpl.Execute(os.Stdout, state)
}

func Test2() {
	text := "{{range .}}- {{.}}\n{{end}}"
	tmpl := template.New("list")
	tmpl = template.Must(tmpl.Parse(text))
	list := []string{"Купить молоко", "Погладить кота", "Вынести мусор"}
	tmpl.Execute(os.Stdout, list)
}
