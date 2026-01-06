package render

import (
	"html/template"
	"time"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"formatDate": formatDate,
		"now":        time.Now,
	}
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
