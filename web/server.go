package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"html/template"

	"github.com/lucasreed/cyoa/book"
)

var story book.Story
var templates map[string]*template.Template

func renderTemplate(w http.ResponseWriter, tmpl string, section string, a book.Arc) {
	//Set up a new local struct so that we can easily access the Section Title
	d := struct {
		Section   string
		StoryData book.Arc
	}{strings.Title(section), a}

	// Get our root dir and template dir so we can find the template files
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	templDir := fmt.Sprintf("%s/templates/", rootDir)
	t, err := template.ParseFiles(templDir + tmpl + ".html")
	if err != nil {
		fmt.Print(err)
	}

	//Render the template into the http.ResponseWriter
	err = t.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func arcHandler(w http.ResponseWriter, r *http.Request) {
	arcKey := r.URL.Path[1:]
	a := story[arcKey]
	renderTemplate(w, "story", arcKey, a)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

// StartServer creates the web server to handle requests
func StartServer(s book.Story) {
	story = s
	http.HandleFunc("/", homeHandler)
	for key := range s {
		http.HandleFunc(fmt.Sprintf("/%s", key), arcHandler)
	}
	log.Print("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
