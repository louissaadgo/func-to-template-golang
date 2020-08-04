package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fn = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type food struct {
	Name string
	ID   string
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fn).ParseFiles("tpl.gohtml"))
}

func main() {
	burger := food{
		"Burger",
		"Brgr",
	}
	pizza := food{
		"Pizza",
		"Pizz",
	}
	foods := []food{
		burger,
		pizza,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", foods)
	if err != nil {
		log.Fatalln(err)
	}
}
