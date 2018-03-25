package main

import (
	"io/ioutil"
	"log"

	"github.com/lucasreed/cyoa/book"
	"github.com/lucasreed/cyoa/web"
)

func main() {

	fileData, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	s := book.NewStory(fileData)
	web.StartServer(s)
}
