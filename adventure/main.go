package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var story string

func main() {
	flag.StringVar(&story, "story", "./public/gopher.json", "Location of JSON story")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)

	http.Handle("/", r)

	port := ":8080"
	log.Printf("Serving from localhost%v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write(parseJSONStory())
}

func parseJSONStory() []byte {
	r, err := os.Open(story)
	d, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}
	return d
}
