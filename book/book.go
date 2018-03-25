package book

import (
	"encoding/json"
	"log"
)

// Arc is a single part, or page, in the story
type Arc struct {
	Title   string   `json:"title"`
	Text    []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// Story is the group of Arcs that make up the full story
type Story map[string]Arc

// NewStory takes json and parses it into a Story
func NewStory(j []byte) Story {
	var ret Story
	if err := json.Unmarshal(j, &ret); err != nil {
		log.Fatal(err)
	}

	return ret
}

// GetPage returns the Arc object of the current section
func (s Story) GetPage(section string) Arc {
	return s[section]
}
