// https://go.dev/pkg/time/?m=old#pkg-index

package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006")
	//	return t.Format("01-02-2006")
	//  return t.Format(time.Kitchen)
	//  return t.Format("02-01-2006")
	//	return t.Format("02 Jan 06 15:04 MST")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalf("template execution failed: %v", err)
	}
}
