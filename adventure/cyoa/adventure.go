package cyoa

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./public/adventure.html"))
}

// NewHandler returns an http.Handler
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := r.URL.Path
	if arc == "/" {
		arc = "intro"
	} else {
		arc = arc[1:]
	}
	tpl.Execute(w, h.s[arc])
}

// Story is a story
type Story map[string]Chapter

// Chapter is a section of a story with options to move forward in the adventure
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option is the gateway to move forward in the story.
// If option.arc is empty, you have reached the end of the story
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// ParseJSONStory will open the targeted JSON file and return a Story
func ParseJSONStory(filename string) (s Story, err error) {
	r, err := os.Open(filename)
	if err != nil {
		log.Println("Could not decode provided JSON reader")
		return nil, err
	}
	d := json.NewDecoder(r)
	if err = d.Decode(&s); err != nil {
		return nil, err
	}
	return s, nil
}
