package web

import (
	"log"
	"os"
	"strings"

	"github.com/alecthomas/template"
	"github.com/lucasreed/cyoa/book"
)

//GenerateHTML creates valid html output with a template and data
func GenerateHTML(templ string, data book.Story, section string) {
	t, err := template.ParseFiles(templ)
	if err != nil {
		log.Fatal(err)
	}

	d := struct {
		Section   string
		StoryData book.Arc
	}{strings.Title(section), data[section]}

	err = t.Execute(os.Stdout, d)
	if err != nil {
		log.Fatal(err)
	}
}
