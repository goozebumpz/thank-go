package lesson11

import (
	"os"
	"text/template"
)

var templateText = `{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. {{if ge .Balance 100 }}Все в порядке{{else if eq .Balance 0}}Доступ заблокирован{{else if lt .Balance 100 }}Пора пополнить{{end}}.`

type User struct {
	Name    string
	Balance int
}

func Test() {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))
	user := User{"Алиса", 500}
	tpl.Execute(os.Stdout, user)
}
