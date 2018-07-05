package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aricalves/gophercises/adventure/cyoa"
)

var filename string
var story cyoa.Story
var port *int

func main() {
	port = flag.Int("port", 8080, "Port to serve HTTP")
	flag.StringVar(&filename, "story", "./public/gopher.json", "File location of JSON story")
	flag.Parse()

	story, err := cyoa.ParseJSONStory(filename)
	if err != nil {
		log.Fatalln("Could not parse JSON story", err)
	}

	http.Handle("/", cyoa.NewHandler(story))

	log.Printf("Serving from localhost:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
